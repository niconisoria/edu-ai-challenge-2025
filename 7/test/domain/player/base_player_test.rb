require_relative "../../test_helper"

class BasePlayerTest < Minitest::Test
  def setup
    super
    @player = BasePlayer.new(@config)
  end

  def test_initial_state
    assert_instance_of Board, @player.board
    assert_equal @config.num_ships, @player.num_ships
    assert_empty @player.guesses
  end

  def test_place_ships
    @player.place_ships
    assert_equal @config.num_ships, @player.board.ships.size
  end

  def test_record_and_has_guessed
    refute @player.has_guessed?("00")
    @player.record_guess("00")
    assert @player.has_guessed?("00")
  end

  def test_decrement_ships
    n = @player.num_ships
    @player.decrement_ships
    assert_equal n - 1, @player.num_ships
  end

  def test_all_ships_sunk
    @player.instance_variable_set(:@num_ships, 0)
    assert @player.all_ships_sunk?
  end

  def test_remaining_ships
    assert_equal @player.num_ships, @player.remaining_ships
  end

  def test_valid_guess
    assert @player.valid_guess?("00")
    @player.record_guess("00")
    refute @player.valid_guess?("00")
    refute @player.valid_guess?("a1")
  end
end
