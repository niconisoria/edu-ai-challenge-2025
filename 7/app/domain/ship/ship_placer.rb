require_relative "ship"

class ShipPlacer
  def initialize(config)
    @config = config
  end

  def place_ships_randomly(board, number_of_ships)
    ships_placed = 0
    max_attempts = 100  # Prevent infinite loops

    while ships_placed < number_of_ships && max_attempts > 0
      orientation = random_orientation
      start_row, start_col = calculate_start_position(orientation, board)

      if start_row && start_col
        locations = calculate_ship_locations(start_row, start_col, orientation)
        unless collision?(locations, board)
          place_ship(board, locations)
          ships_placed += 1
        end
      end

      max_attempts -= 1
    end

    raise "Failed to place all ships after #{max_attempts} attempts" if ships_placed < number_of_ships
  end

  private

  def random_orientation
    (rand < 0.5) ? "horizontal" : "vertical"
  end

  def calculate_start_position(orientation, board)
    if orientation == "horizontal"
      [rand(@config.board_size), rand(@config.board_size - @config.ship_length + 1)]
    else
      [rand(@config.board_size - @config.ship_length + 1), rand(@config.board_size)]
    end
  end

  def calculate_ship_locations(start_row, start_col, orientation)
    @config.ship_length.times.map do |i|
      {
        row: (orientation == "horizontal") ? start_row : start_row + i,
        col: (orientation == "horizontal") ? start_col + i : start_col
      }
    end
  end

  def collision?(locations, board)
    locations.any? do |loc|
      loc[:row] >= @config.board_size ||
        loc[:col] >= @config.board_size ||
        !board.get_cell(loc[:row], loc[:col]).empty?
    end
  end

  def place_ship(board, locations)
    new_ship = Ship.new
    locations.each do |loc|
      location_str = "#{loc[:row]}#{loc[:col]}"
      new_ship.add_location(location_str)
      board.get_cell(loc[:row], loc[:col]).place_ship
    end
    board.ships << new_ship
  end
end
