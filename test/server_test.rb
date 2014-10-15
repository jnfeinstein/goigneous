#!/usr/bin/ruby

require 'rest_client'
require 'json'

def make_request(method, url, payload=nil)
  RestClient::Request.execute(:method => method,
                              :url => url,
                              :payload => payload,
                              :user => "igneous",
                              :password => "joel")
end

def check_http_success(response)
  if response.code != 200
    puts "#{response.code}: #{response.to_str}"
    exit
  end
end

def check_url(url, expected)
  if url != expected
    puts "Url was \"#{url}\", expected \"#{expected}\""
    exit
  end
end

def check_content(content, expected)
  if content != expected
    puts "Content was \"#{content}\", expected \"#{expected}\""
    exit
  end
end

def check_error_type(type, expected)
  if type != expected
    puts "Type was \"#{type}\", expected \"#{expected}\""
    exit
  end
end

def check_time(started, ended, expected)
  if ended - started > expected
    puts "Took #{ended - started}s, should have been less than #{expected}s"
    exit
  end
end

host = ARGV.length > 0 && ARGV[0] == "-p" ? 'igneous.joelf.me' : 'localhost:3000'
puts "Using host \"#{host}\""

content = {:abc => "123"}
update_content = {:xyz => "123"}

puts "Adding new entry"
response = make_request(:post, "http://#{host}/documents/new", {:content => content.to_json})
check_http_success(response)

url = response.to_str
puts "Verifying content"
response = make_request(:get, url)
check_http_success(response)
check_content(response.to_str, content.to_json)

puts "Updating content"
response = make_request(:put, url, {:content => update_content.to_json})
check_http_success(response)
check_url(response.to_str, url)

puts "Verifying updated content"
response = make_request(:get, url)
check_http_success(response)
check_content(response.to_str, update_content.to_json)

puts "Deleting content"
response = make_request(:delete, url)
check_http_success(response)

puts "Ensuring that content is protected"
begin
  RestClient.get("http://#{host}/documents/3")
  exit
rescue Exception => e
  check_error_type(e.class, RestClient::Unauthorized)
end

puts "Checking multiple connections"
url = make_request(:post, "http://#{host}/documents/new", {:content => content.to_json}).to_str
started = Time.now
threads = 5.times.map do
  Thread.new do
    response = make_request(:get, url)
    check_http_success(response)
    check_content(response.to_str, content.to_json)
    sleep 9
  end
end
threads.each{ |t| t.join }
check_time(started, Time.now, 5)

puts "Success"
