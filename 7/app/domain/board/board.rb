require_relative "cell"
require_relative "../ship/ship_placer"
require_relative "../../infrastructure/display/board_formatter"

class Board
  attr_reader :cells, :ships, :config

  def initialize(config)
    @config = config
    @cells = []
    @ships = []
    @ship_placer = ShipPlacer.new(config)
    @formatter = BoardFormatter.new(config)
    create_grid
  end

  def create_grid
    @cells = Array.new(@config.board_size) do
      Array.new(@config.board_size) { Cell.new }
    end
  end

  def set_cell(row, col, value)
    case value
    when "S"
      @cells[row][col].place_ship
    when "X"
      @cells[row][col].hit
    when "O"
      @cells[row][col].miss
    else
      @cells[row][col] = Cell.new(value)
    end
  end

  def get_cell(row, col)
    @cells[row][col]
  end

  def valid_position?(row, col)
    @config.valid_position?(row, col)
  end

  def print_row(row_num, hide_ships = false)
    @formatter.format_row(row_num, @cells, hide_ships)
  end

  def place_ships_randomly(number_of_ships)
    @ship_placer.place_ships_randomly(self, number_of_ships)
  end

  def empty_cell?(row, col)
    get_cell(row, col).empty?
  end

  def hit_cell?(row, col)
    get_cell(row, col).hit?
  end

  def miss_cell?(row, col)
    get_cell(row, col).miss?
  end
end
