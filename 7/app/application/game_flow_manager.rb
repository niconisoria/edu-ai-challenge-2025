require_relative "../domain/game/game_state"
require_relative "../infrastructure/display/game_display"
require_relative "../infrastructure/input/game_input"
require_relative "../application/turn_manager"

class GameFlowManager
  def initialize(state, display, input, turn_manager)
    @state = state
    @display = display
    @input = input
    @turn_manager = turn_manager
  end

  def start_game
    @state.start_game
    @display.print_game_start(@state.cpu.num_ships)
    game_loop
  end

  def game_loop
    print_boards

    if @state.check_game_over
      @display.print_game_over(@state.player_won?)
      return
    end

    @display.print_turn_prompt
    guess = @input.get_player_guess

    if @turn_manager.process_player_turn(guess)
      @turn_manager.process_cpu_turn
    end

    game_loop
  end

  private

  def print_boards
    @display.print_boards(@state.player.board, @state.cpu.board)
  end
end
