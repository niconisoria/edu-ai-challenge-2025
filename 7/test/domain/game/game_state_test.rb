require_relative "../../test_helper"

class GameStateTest < Minitest::Test
  def setup
    super
    @state = GameState.new(@config)
  end

  def test_initial_state
    assert_instance_of HumanPlayer, @state.player
    assert_instance_of CPUPlayer, @state.cpu
    refute @state.is_game_over
    assert_nil @state.winner
  end

  def test_start_game_places_ships
    @state.start_game
    assert @state.player.board.ships.any?
    assert @state.cpu.board.ships.any?
  end

  def test_check_game_over
    @state.start_game
    @state.cpu.instance_variable_set(:@num_ships, 0)
    assert @state.check_game_over
    assert @state.is_game_over
    assert_equal :player, @state.winner
    assert @state.player_won?
    refute @state.cpu_won?

    @state = GameState.new(@config)
    @state.start_game
    @state.player.instance_variable_set(:@num_ships, 0)
    assert @state.check_game_over
    assert @state.is_game_over
    assert_equal :cpu, @state.winner
    refute @state.player_won?
    assert @state.cpu_won?
  end
end
