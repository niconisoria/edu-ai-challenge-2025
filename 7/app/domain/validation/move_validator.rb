class MoveValidator
  HIT_RESULTS = {
    already_hit: ->(board, row, col, ship, player_type, messages) {
      messages.format_already_hit
      :already_hit
    },
    sunk: ->(board, row, col, ship, player_type, messages) {
      board.set_cell(row, col, "X")
      messages.format_hit(player_type, "#{row}#{col}")
      messages.format_ship_sunk(player_type) if player_type == "CPU"
      :sunk
    },
    hit: ->(board, row, col, ship, player_type, messages) {
      board.set_cell(row, col, "X")
      messages.format_hit(player_type, "#{row}#{col}")
      :hit
    },
    miss: ->(board, row, col, ship, player_type, messages) {
      board.set_cell(row, col, "O")
      :miss
    }
  }.freeze

  def initialize(messages)
    @messages = messages
  end

  def process_hit(board, row, col, ship, player_type)
    location = "#{row}#{col}"

    return HIT_RESULTS[:already_hit].call(board, row, col, ship, player_type, @messages) if ship.hit_at?(location)

    if ship.record_hit(location)
      return HIT_RESULTS[:sunk].call(board, row, col, ship, player_type, @messages) if ship.sunk?
      return HIT_RESULTS[:hit].call(board, row, col, ship, player_type, @messages)
    end

    HIT_RESULTS[:miss].call(board, row, col, ship, player_type, @messages)
  end

  def valid_move?(row, col, board, guesses)
    board.valid_position?(row, col) && !guesses.include?("#{row}#{col}")
  end
end
