require_relative "../../test_helper"

class ShipPlacerTest < Minitest::Test
  class TestConfig
    attr_reader :board_size, :ship_length
    def initialize
      @board_size = 5
      @ship_length = 2
    end
  end

  def setup
    @config = TestConfig.new
    @board = Board.new(@config)
    @placer = ShipPlacer.new(@config)
  end

  def test_place_ships_randomly
    @placer.place_ships_randomly(@board, 2)
    assert_equal 2, @board.ships.size
    @board.ships.each do |ship|
      ship.locations.each do |loc|
        row = loc[0].to_i
        col = loc[1].to_i
        assert @board.get_cell(row, col).ship?
      end
    end
  end
end
