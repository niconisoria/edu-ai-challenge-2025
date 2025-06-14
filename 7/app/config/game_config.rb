class GameConfig
  attr_reader :board_size, :num_ships, :ship_length

  def initialize(board_size: 10, num_ships: 3, ship_length: 3)
    @board_size = board_size
    @num_ships = num_ships
    @ship_length = ship_length
  end

  def valid_position?(row, col)
    row >= 0 && row < @board_size && col >= 0 && col < @board_size
  end
end
