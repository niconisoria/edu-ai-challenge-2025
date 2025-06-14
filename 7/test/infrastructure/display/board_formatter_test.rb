require_relative "../../test_helper"

class BoardFormatterTest < Minitest::Test
  def setup
    super
    @formatter = BoardFormatter.new(@config)
    @board = Board.new(@config)
  end

  def test_initialization
    assert_instance_of BoardFormatter, @formatter
    assert_same @config, @formatter.instance_variable_get(:@config)
  end

  def test_format_header
    expected = "  0 1 2 3 4 5 6 7 8 9"
    assert_equal expected, @formatter.format_header
  end

  def test_format_row_with_hidden_ships
    @board.set_cell(0, 0, "S")
    @board.set_cell(0, 1, "X")
    @board.set_cell(0, 2, "O")
    
    expected = "0 ~ X O ~ ~ ~ ~ ~ ~ ~ "
    assert_equal expected, @formatter.format_row(0, @board.cells, true)
  end

  def test_format_row_without_hiding_ships
    @board.set_cell(0, 0, "S")
    @board.set_cell(0, 1, "X")
    @board.set_cell(0, 2, "O")
    
    expected = "0 S X O ~ ~ ~ ~ ~ ~ ~ "
    assert_equal expected, @formatter.format_row(0, @board.cells, false)
  end

  def test_format_boards
    player_board = Board.new(@config)
    cpu_board = Board.new(@config)

    # Set up some cells
    player_board.set_cell(0, 0, "S")
    player_board.set_cell(0, 1, "X")
    player_board.set_cell(0, 2, "O")

    cpu_board.set_cell(1, 1, "S")
    cpu_board.set_cell(1, 2, "X")
    cpu_board.set_cell(1, 3, "O")

    expected = [
      "\n   --- OPPONENT BOARD ---          --- YOUR BOARD ---",
      "  0 1 2 3 4 5 6 7 8 9       0 1 2 3 4 5 6 7 8 9",
      "0 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~     0 S X O ~ ~ ~ ~ ~ ~ ~ ",
      "1 ~ ~ X O ~ ~ ~ ~ ~ ~     1 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ",
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

    assert_equal expected, @formatter.format_boards(player_board, cpu_board)
  end
end 