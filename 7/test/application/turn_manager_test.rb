require_relative "../test_helper"

class TurnManagerTest < Minitest::Test
  def setup
    super
    @state = GameState.new(@config)
    @messages = GameMessages.new(@config)
    @turn_manager = TurnManager.new(@state, @config, @messages)
  end

  def test_initialization
    assert_instance_of TurnManager, @turn_manager
    assert_same @state, @turn_manager.instance_variable_get(:@state)
    assert_same @config, @turn_manager.instance_variable_get(:@config)
    assert_same @messages, @turn_manager.instance_variable_get(:@messages)
  end

  def test_current_player_returns_correct_player
    assert_equal @state.player, @turn_manager.current_player
    @turn_manager.switch_turn
    assert_equal @state.cpu, @turn_manager.current_player
  end

  def test_switch_turn_alternates_players
    assert_equal @state.player, @turn_manager.current_player
    @turn_manager.switch_turn
    assert_equal @state.cpu, @turn_manager.current_player
    @turn_manager.switch_turn
    assert_equal @state.player, @turn_manager.current_player
  end
end
