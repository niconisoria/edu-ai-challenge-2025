require_relative "../messaging/game_messages"
require_relative "../display/board_formatter"

class GameDisplay
  def initialize(config)
    @config = config
    @messages = GameMessages.new(config)
    @formatter = BoardFormatter.new(config)
  end

  def print_boards(player_board, cpu_board)
    @formatter.format_boards(player_board, cpu_board).each { |line| puts line }
  end

  def print_game_over(player_won)
    puts @messages.format_game_over(player_won)
  end

  def print_game_start(num_ships)
    @messages.format_game_start(num_ships).each { |msg| puts msg }
  end

  def print_turn_prompt
    print @messages.format_turn_prompt
  end
end
