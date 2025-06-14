require_relative "../domain/player/cpu_strategy_manager"
require_relative "../infrastructure/input/input_processor"
require_relative "../infrastructure/messaging/game_messages"
require_relative "../domain/validation/move_validator"
require_relative "../domain/game/game_turn"

class TurnManager
  def initialize(state, config, messages)
    @state = state
    @config = config
    @messages = messages
    @player_turn = GameTurn.new(state.player, state.cpu)
    @cpu_turn = GameTurn.new(state.cpu, state.player)
    @cpu_strategy = CPUStrategyManager.new(config)
    @move_validator = MoveValidator.new(messages)
    @turn = :player
  end

  def process_player_turn(guess)
    return false unless guess
    @player_turn.process_guess(guess)
  end

  def process_cpu_turn
    puts @messages.format_cpu_turn
    guess_str = get_valid_cpu_guess
    return unless guess_str

    puts @messages.format_cpu_target(guess_str)
    @cpu_turn.process_guess(guess_str)
  end

  def current_player
    (@turn == :player) ? @state.player : @state.cpu
  end

  def switch_turn
    @turn = (@turn == :player) ? :cpu : :player
  end

  private

  def get_valid_cpu_guess
    guess_str = nil
    until guess_str
      guess_str = @cpu_strategy.next_guess(@state.cpu.guesses)
    end
    guess_str
  end

  def process_cpu_guess(guess_str)
    guess_row = guess_str[0].to_i
    guess_col = guess_str[1].to_i
    puts @messages.format_cpu_target(guess_str)
    @state.cpu.record_guess(guess_str)

    hit_ship = find_hit_ship(guess_str)
    if hit_ship
      process_hit(guess_row, guess_col, hit_ship)
    else
      process_miss(guess_row, guess_col)
    end
  end

  def find_hit_ship(guess_str)
    @state.player.board.ships.find { |ship| ship.locations.include?(guess_str) }
  end

  def process_hit(guess_row, guess_col, ship)
    result = @move_validator.process_hit(@state.player.board, guess_row, guess_col, ship, "CPU")

    if result == :sunk
      @state.player.decrement_ships
      @cpu_strategy.record_sunk
    elsif result == :hit
      @cpu_strategy.record_hit(guess_row, guess_col, @state.cpu.guesses)
    end
  end

  def process_miss(guess_row, guess_col)
    @state.player.board.set_cell(guess_row, guess_col, "O")
    puts @messages.format_miss("CPU", "#{guess_row}#{guess_col}")
  end
end
