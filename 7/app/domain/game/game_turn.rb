require_relative "../../infrastructure/input/input_processor"
require_relative "../../infrastructure/messaging/game_messages"
require_relative "../validation/move_validator"

class GameTurn
  attr_reader :current_player, :opponent

  def initialize(player, opponent)
    @player = player
    @opponent = opponent
    @messages = GameMessages.new(opponent.board.config)
    @processor = InputProcessor.new(opponent.board.config, @messages)
    @move_validator = MoveValidator.new(@messages)
    @player_type = player.is_a?(CPUPlayer) ? "CPU" : "PLAYER"
  end

  def process_guess(guess)
    parsed_guess = @processor.parse_guess(guess)
    return false unless parsed_guess

    if @player.has_guessed?(parsed_guess[:string])
      puts @messages.format_already_guessed
      return false
    end
    @player.record_guess(parsed_guess[:string])

    hit = false
    @opponent.board.ships.each do |ship|
      if ship.locations.include?(parsed_guess[:string])
        if ship.hit_at?(parsed_guess[:string])
          puts @messages.format_already_hit
          hit = true
          break
        end

        if ship.record_hit(parsed_guess[:string])
          @opponent.board.set_cell(parsed_guess[:row], parsed_guess[:col], "X")
          puts @messages.format_hit(@player_type, parsed_guess[:string])
          hit = true

          if ship.sunk?
            @opponent.decrement_ships
          end
          break
        end
      end
    end

    unless hit
      @opponent.board.set_cell(parsed_guess[:row], parsed_guess[:col], "O")
      puts @messages.format_miss(@player_type, parsed_guess[:string])
    end

    true
  end

  def valid_and_new_guess?(row, col, guess_list)
    return false unless @opponent.board.valid_position?(row, col)
    guess_str = "#{row}#{col}"
    !guess_list.include?(guess_str)
  end
end
