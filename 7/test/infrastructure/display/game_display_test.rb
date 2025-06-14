require_relative "../../test_helper"

class GameDisplayTest < Minitest::Test
  def setup
    super
    @display = GameDisplay.new(@config)
    @player_board = Board.new(@config)
    @cpu_board = Board.new(@config)
  end

  def test_initialization
    assert_instance_of GameDisplay, @display
    assert_instance_of GameMessages, @display.instance_variable_get(:@messages)
    assert_instance_of BoardFormatter, @display.instance_variable_get(:@formatter)
  end

  def test_print_boards
    expected_output = [
      "\n   --- OPPONENT BOARD ---          --- YOUR BOARD ---",
      "  0 1 2 3 4 5 6 7 8 9       0 1 2 3 4 5 6 7 8 9",
      "0 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~     0 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ",
      "1 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~     1 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ",
      "2 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~     2 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ",
      "3 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~     3 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ",
      "4 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~     4 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ",
      "5 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~     5 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ",
      "6 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~     6 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ",
      "7 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~     7 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ",
      "8 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~     8 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ",
      "9 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~     9 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ",
      "\n"
    ]

    assert_output(expected_output.join("\n")) do
      @display.print_boards(@player_board, @cpu_board)
    end
  end

  def test_print_game_over_player_won
    expected_output = "\n*** CONGRATULATIONS! You sunk all enemy battleships! ***\n"
    assert_output(expected_output) do
      @display.print_game_over(true)
    end
  end

  def test_print_game_over_cpu_won
    expected_output = "\n*** GAME OVER! The CPU sunk all your battleships! ***\n"
    assert_output(expected_output) do
      @display.print_game_over(false)
    end
  end

  def test_print_game_start
    expected_output = [
      "\nLet's play Sea Battle!",
      "Try to sink the 3 enemy ships."
    ].join("\n") + "\n"

    assert_output(expected_output) do
      @display.print_game_start(3)
    end
  end

  def test_print_turn_prompt
    expected_output = "Enter your guess (e.g., 00): "
    assert_output(expected_output) do
      @display.print_turn_prompt
    end
  end
end 