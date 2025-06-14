require_relative "../test_helper"

class GameFlowManagerTest < Minitest::Test
  def setup
    super
    @state = GameState.new(@config)
    @display = GameDisplay.new(@config)
    @input = GameInput.new(@config)
    @turn_manager = TurnManager.new(@state, @config, GameMessages.new(@config))
    @flow_manager = GameFlowManager.new(@state, @display, @input, @turn_manager)
  end

  def test_initialization
    assert_instance_of GameFlowManager, @flow_manager
    assert_same @state, @flow_manager.instance_variable_get(:@state)
    assert_same @display, @flow_manager.instance_variable_get(:@display)
    assert_same @input, @flow_manager.instance_variable_get(:@input)
    assert_same @turn_manager, @flow_manager.instance_variable_get(:@turn_manager)
  end
end
