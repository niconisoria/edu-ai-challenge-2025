require "simplecov"
SimpleCov.start do
  add_filter "/test/"
end

require "minitest/autorun"
require "minitest/pride"

# Add the app directory to the load path
$LOAD_PATH.unshift(File.expand_path("../app", __dir__))

# Require all necessary files
require "domain/board/board"
require "domain/board/cell"
require "domain/board/position"
require "domain/game/game_state"
require "domain/game/game_turn"
require "domain/player/base_player"
require "domain/player/human_player"
require "domain/player/cpu_player"
require "domain/player/cpu_strategy_manager"
require "domain/ship/ship"
require "domain/ship/ship_placer"
require "domain/validation/move_validator"
require "infrastructure/display/board_formatter"
require "infrastructure/display/game_display"
require "infrastructure/input/game_input"
require "infrastructure/input/input_processor"
require "infrastructure/messaging/game_messages"
require "application/game_controller"
require "application/game_flow_manager"
require "application/turn_manager"
require "config/game_config"

class Minitest::Test
  def setup
    @config = GameConfig.new
  end
end
