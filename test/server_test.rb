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
  if response.code == 200
    true
  else
    puts "#{response.code}: #{response.to_str}"
    false
  end
end

def check_url(url, expected)
  if url != expected
    puts "Url was \"#{url}\", expected \"#{expected}\""
    false
  else
    true
  end
end

def check_content(content, expected)
  if content != expected
    puts "Content was \"#{content}\", expected \"#{expected}\""
    false
  else
    true
  end
end

host = ARGV.length > 0 && ARGV[0] == "-p" ? 'goigneous.herokuapp.com' : 'localhost:3000'
puts "Using host \"#{host}\""

content = {:abc => "123"}
update_content = {:xyz => "123"}

puts "Adding new entry"
response = make_request(:post, "http://#{host}/documents/new", {:content => content.to_json})
exit unless check_http_success(response)

url = response.to_str
puts "Verifying content"
response = make_request(:get, url)
exit unless check_http_success(response)
exit unless check_content(response.to_str, content.to_json)

puts "Updating content"
response = make_request(:put, url, {:content => update_content.to_json})
exit unless check_http_success(response)
exit unless check_url(response.to_str, url)

puts "Verifying updated content"
response = make_request(:get, url)
exit unless check_http_success(response)
exit unless check_content(response.to_str, update_content.to_json)

puts "Deleting content"
response = make_request(:delete, url)
exit unless check_http_success(response)

puts "Success"
