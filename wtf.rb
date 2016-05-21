#!/usr/bin/env ruby

def battleship(v)
  l = 'MISS'
  l = 'HIT' if v
  "#{l} : #{v}"
end

def main
 [0, '', [], nil].each do |v|
    puts battleship(v)
  end
end

main
