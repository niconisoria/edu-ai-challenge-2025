require_relative "../test_helper"

class GameControllerTest < Minitest::Test
  def setup
    super
    @game_controller = GameController.new
  end

  def test_initialization
    assert_instance_of GameController, @game_controller
    assert_instance_of GameConfig, @game_controller.instance_variable_get(:@config)
    assert_instance_of GameState, @game_controller.instance_variable_get(:@state)
    assert_instance_of GameMessages, @game_controller.instance_variable_get(:@messages)
    assert_instance_of GameDisplay, @game_controller.instance_variable_get(:@display)
    assert_instance_of GameInput, @game_controller.instance_variable_get(:@input)
    assert_instance_of TurnManager, @game_controller.instance_variable_get(:@turn_manager)
    assert_instance_of GameFlowManager, @game_controller.instance_variable_get(:@flow_manager)
  end

  def test_start_game_calls_flow_manager
    flow_manager = Minitest::Mock.new
    flow_manager.expect :start_game, nil

    @game_controller.instance_variable_set(:@flow_manager, flow_manager)
    @game_controller.start_game

    assert_mock flow_manager
  end
end
