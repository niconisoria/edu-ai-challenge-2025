require_relative "../../test_helper"

class GameInputTest < Minitest::Test
  def setup
    super
    @input = GameInput.new(@config)
  end

  def test_initialization
    assert_instance_of GameInput, @input
    assert_instance_of GameMessages, @input.instance_variable_get(:@messages)
    assert_instance_of InputProcessor, @input.instance_variable_get(:@processor)
  end

  def test_get_player_guess
    input = "12"
    $stdin = StringIO.new(input)
    
    assert_equal "12", @input.get_player_guess
  ensure
    $stdin = STDIN
  end

  def test_parse_guess_valid
    result = @input.parse_guess("12")
    assert_equal 1, result[:row]
    assert_equal 2, result[:col]
    assert_equal "12", result[:string]
  end

  def test_parse_guess_invalid
    assert_nil @input.parse_guess("invalid")
  end
end 