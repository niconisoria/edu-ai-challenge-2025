require_relative "../messaging/game_messages"
require_relative "input_processor"

class GameInput
  def initialize(config)
    @config = config
    @messages = GameMessages.new(config)
    @processor = InputProcessor.new(config, @messages)
  end

  def get_player_guess
    guess = gets.chomp
    @processor.process_guess(guess)
  end

  def parse_guess(guess)
    return nil unless guess && guess.match?(/^\d{2}$/)
    @processor.parse_guess(guess)
  end
end
