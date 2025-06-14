require_relative "../../test_helper"

class BoardTest < Minitest::Test
  def setup
    super
    @board = Board.new(@config)
  end

  def test_initialization
    assert_equal @config, @board.config
    assert_empty @board.ships
    assert_equal @config.board_size, @board.cells.length
    @board.cells.each do |row|
      assert_equal @config.board_size, row.length
      row.each { |cell| assert_instance_of Cell, cell }
    end
  end

  def test_set_cell
    @board.set_cell(0, 0, "S")
    assert_equal "S", @board.get_cell(0, 0).value

    @board.set_cell(1, 1, "X")
    assert_equal "X", @board.get_cell(1, 1).value

    @board.set_cell(2, 2, "O")
    assert_equal "O", @board.get_cell(2, 2).value
  end

  def test_valid_position
    assert @board.valid_position?(0, 0)
    assert @board.valid_position?(@config.board_size - 1, @config.board_size - 1)
    refute @board.valid_position?(-1, 0)
    refute @board.valid_position?(0, -1)
    refute @board.valid_position?(@config.board_size, 0)
    refute @board.valid_position?(0, @config.board_size)
  end

  def test_empty_cell
    assert @board.empty_cell?(0, 0)
    @board.set_cell(0, 0, "S")
    refute @board.empty_cell?(0, 0)
  end

  def test_hit_cell
    refute @board.hit_cell?(0, 0)
    @board.set_cell(0, 0, "X")
    assert @board.hit_cell?(0, 0)
  end

  def test_miss_cell
    refute @board.miss_cell?(0, 0)
    @board.set_cell(0, 0, "O")
    assert @board.miss_cell?(0, 0)
  end

  def test_place_ships_randomly
    num_ships = 3
    @board.place_ships_randomly(num_ships)
    assert_equal num_ships, @board.ships.length

    # Verify ships are placed correctly
    ship_cells = @board.ships.flat_map(&:locations)
    ship_cells.each do |location|
      row = location[0].to_i
      col = location[1].to_i
      assert @board.get_cell(row, col).ship?
    end
  end
end
