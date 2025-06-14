require_relative "../../test_helper"

class GameMessagesTest < Minitest::Test
  def setup
    super
    @messages = GameMessages.new(@config)
  end

  def test_initialization
    assert_instance_of GameMessages, @messages
    assert_same @config, @messages.instance_variable_get(:@config)
    assert_instance_of Hash, @messages.instance_variable_get(:@templates)
  end

  def test_format_board_header
    expected = "  0 1 2 3 4 5 6 7 8 9 "
    assert_equal expected, @messages.format_board_header
  end

  def test_format_game_over_player_won
    expected = "\n*** CONGRATULATIONS! You sunk all enemy battleships! ***"
    assert_equal expected, @messages.format_game_over(true)
  end

  def test_format_game_over_cpu_won
    expected = "\n*** GAME OVER! The CPU sunk all your battleships! ***"
    assert_equal expected, @messages.format_game_over(false)
  end

  def test_format_game_start
    expected = [
      "\nLet's play Sea Battle!",
      "Try to sink the 3 enemy ships."
    ]
    assert_equal expected, @messages.format_game_start(3)
  end

  def test_format_turn_prompt
    expected = "Enter your guess (e.g., 00): "
    assert_equal expected, @messages.format_turn_prompt
  end

  def test_format_invalid_guess_format
    expected = "Oops, input must be exactly two digits (e.g., 00, 34, 98)."
    assert_equal expected, @messages.format_invalid_guess_format
  end

  def test_format_invalid_guess_position
    expected = "Oops, please enter valid row and column numbers between 0 and 9."
    assert_equal expected, @messages.format_invalid_guess_position
  end

  def test_format_cpu_turn
    expected = "\n--- CPU's Turn ---"
    assert_equal expected, @messages.format_cpu_turn
  end

  def test_format_cpu_target
    expected = "CPU targets: 12"
    assert_equal expected, @messages.format_cpu_target("12")
  end

  def test_format_hit
    expected = "Player HIT at 12!"
    assert_equal expected, @messages.format_hit("Player", "12")
  end

  def test_format_miss
    expected = "Player MISS at 12."
    assert_equal expected, @messages.format_miss("Player", "12")
  end

  def test_format_ship_sunk
    expected = "Player sunk your battleship!"
    assert_equal expected, @messages.format_ship_sunk("Player")
  end

  def test_format_already_hit
    expected = "You already hit that spot!"
    assert_equal expected, @messages.format_already_hit
  end

  def test_format_already_guessed
    expected = "You already guessed that location!"
    assert_equal expected, @messages.format_already_guessed
  end
end 