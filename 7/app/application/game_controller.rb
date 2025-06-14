require_relative "../domain/game/game_state"
require_relative "../infrastructure/display/game_display"
require_relative "../infrastructure/input/game_input"
require_relative "turn_manager"
require_relative "game_flow_manager"
require_relative "../infrastructure/messaging/game_messages"
require_relative "../config/game_config"

class GameController
  def initialize
    @config = GameConfig.new
    @state = GameState.new(@config)
    @messages = GameMessages.new(@config)
    @display = GameDisplay.new(@config)
    @input = GameInput.new(@config)
    @turn_manager = TurnManager.new(@state, @config, @messages)
    @flow_manager = GameFlowManager.new(@state, @display, @input, @turn_manager)
  end

  def start_game
    @flow_manager.start_game
  end
end
