require_relative "../../test_helper"

class InputProcessorTest < Minitest::Test
  def setup
    super
    @config = GameConfig.new(board_size: 5)
    @messages = GameMessages.new(@config)
    @processor = InputProcessor.new(@config, @messages)
  end

  def test_initialization
    assert_instance_of InputProcessor, @processor
    assert_same @config, @processor.instance_variable_get(:@config)
    assert_same @messages, @processor.instance_variable_get(:@messages)
  end

  def test_process_guess_valid
    assert_output("") do
      assert_equal "12", @processor.process_guess("12")
    end
  end

  def test_process_guess_invalid_format
    assert_output(@messages.format_invalid_guess_format + "\n") do
      assert_nil @processor.process_guess("123")
    end
  end

  def test_process_guess_invalid_position
    assert_output(@messages.format_invalid_guess_position + "\n") do
      assert_nil @processor.process_guess("99")
    end
  end

  def test_valid_guess_format_valid
    assert_output("") do
      assert @processor.valid_guess_format?("12")
    end
  end

  def test_valid_guess_format_invalid
    assert_output(@messages.format_invalid_guess_format + "\n") do
      refute @processor.valid_guess_format?("123")
    end
  end

  def test_valid_guess_position_valid
    assert_output("") do
      assert @processor.valid_guess_position?("12")
    end
  end

  def test_valid_guess_position_invalid
    assert_output(@messages.format_invalid_guess_position + "\n") do
      refute @processor.valid_guess_position?("99")
    end
  end

  def test_parse_guess_valid
    result = @processor.parse_guess("12")
    assert_equal 1, result[:row]
    assert_equal 2, result[:col]
    assert_equal "12", result[:string]
  end

  def test_parse_guess_invalid
    assert_nil @processor.parse_guess(nil)
  end
end 