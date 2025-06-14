class GameMessages
  def initialize(config)
    @config = config
    @templates = {
      board_header: "  %s",
      game_over: {
        player: "\n*** CONGRATULATIONS! You sunk all enemy battleships! ***",
        cpu: "\n*** GAME OVER! The CPU sunk all your battleships! ***"
      },
      game_start: [
        "\nLet's play Sea Battle!",
        "Try to sink the %d enemy ships."
      ],
      turn_prompt: "Enter your guess (e.g., 00): ",
      invalid_guess: {
        format: "Oops, input must be exactly two digits (e.g., 00, 34, 98).",
        position: "Oops, please enter valid row and column numbers between 0 and %d."
      },
      cpu_turn: "\n--- CPU's Turn ---",
      cpu_target: "CPU targets: %s",
      hit: "%s HIT at %s!",
      miss: "%s MISS at %s.",
      ship_sunk: "%s sunk your battleship!",
      already_hit: "You already hit that spot!",
      already_guessed: "You already guessed that location!"
    }
  end

  def format_board_header
    header = "  "
    @config.board_size.times { |h| header += "#{h} " }
    header
  end

  def format_game_over(player_won)
    @templates[:game_over][player_won ? :player : :cpu]
  end

  def format_game_start(num_ships)
    @templates[:game_start].map { |msg| msg % num_ships }
  end

  def format_turn_prompt
    @templates[:turn_prompt]
  end

  def format_invalid_guess_format
    @templates[:invalid_guess][:format]
  end

  def format_invalid_guess_position
    @templates[:invalid_guess][:position] % (@config.board_size - 1)
  end

  def format_cpu_turn
    @templates[:cpu_turn]
  end

  def format_cpu_target(guess)
    @templates[:cpu_target] % guess
  end

  def format_hit(player_type, location = "")
    @templates[:hit] % [player_type, location]
  end

  def format_miss(player_type, location = "")
    @templates[:miss] % [player_type, location]
  end

  def format_ship_sunk(player_type)
    @templates[:ship_sunk] % player_type
  end

  def format_already_hit
    @templates[:already_hit]
  end

  def format_already_guessed
    @templates[:already_guessed]
  end
end
