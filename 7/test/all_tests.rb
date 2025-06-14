require_relative "test_helper"

Dir.glob(File.join(__dir__, "**", "*_test.rb")).sort.each do |file|
  require file unless file == __FILE__
end
