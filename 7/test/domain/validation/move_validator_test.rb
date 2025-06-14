require_relative "../../test_helper"

class MoveValidatorTest < Minitest::Test
  def setup
    super
    @messages = GameMessages.new(@config)
    @validator = MoveValidator.new(@messages)
    @board = Board.new(@config)
    @ship = Ship.new
    @ship.add_location("00")
    @ship.add_location("01")
    @board.set_cell(0, 0, "S")
    @board.set_cell(0, 1, "S")
  end

  def test_valid_move
    guesses = []
    assert @validator.valid_move?(0, 0, @board, guesses)
    guesses << "00"
    refute @validator.valid_move?(0, 0, @board, guesses)
    refute @validator.valid_move?(-1, 0, @board, guesses)
  end

  def test_process_hit_already_hit
    @ship.record_hit("00")
    assert_equal :already_hit, @validator.process_hit(@board, 0, 0, @ship, "PLAYER")
  end

  def test_process_hit_sunk
    @validator.process_hit(@board, 0, 0, @ship, "PLAYER")
    result = @validator.process_hit(@board, 0, 1, @ship, "PLAYER")
    assert_equal :sunk, result
  end

  def test_process_hit_hit
    result = @validator.process_hit(@board, 0, 0, @ship, "PLAYER")
    assert_equal :hit, result
  end

  def test_process_hit_miss
    result = @validator.process_hit(@board, 1, 1, @ship, "PLAYER")
    assert_equal :miss, result
  end
end
