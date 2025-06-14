require_relative "../player/human_player"
require_relative "../player/cpu_player"
require_relative "../player/cpu_strategy_manager"

class GameState
  attr_reader :player, :cpu, :is_game_over, :winner

  def initialize(config)
    @config = config
    @player = HumanPlayer.new(config)
    @cpu = CPUPlayer.new(config, CPUStrategyManager.new(config))
    @is_game_over = false
    @winner = nil
  end

  def start_game
    @player.place_ships
    @cpu.place_ships
  end

  def check_game_over
    return false unless @player.num_ships.zero? || @cpu.num_ships.zero?

    @is_game_over = true
    @winner = @cpu.num_ships.zero? ? :player : :cpu
    true
  end

  def player_won?
    @winner == :player
  end

  def cpu_won?
    @winner == :cpu
  end
end
