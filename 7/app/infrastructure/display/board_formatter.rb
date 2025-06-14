class BoardFormatter
  CELL_STATES = {
    "S" => "~",  # Ship (hidden)
    "X" => "X",  # Hit
    "O" => "O"   # Miss
  }.freeze

  def initialize(config)
    @config = config
  end

  def format_row(row_num, cells, hide_ships = false)
    row_str = "#{row_num} "
    @config.board_size.times do |j|
      cell = cells[row_num][j]
      value = cell.to_s
      value = CELL_STATES[value] if hide_ships && CELL_STATES.key?(value)
      row_str += "#{value} "
    end
    row_str
  end

  def format_header
    "  " + (0...@config.board_size).map(&:to_s).join(" ")
  end

  def format_boards(player_board, cpu_board)
    [
      "\n   --- OPPONENT BOARD ---          --- YOUR BOARD ---",
      "#{format_header}     #{format_header}",
      *@config.board_size.times.map { |i|
        "#{format_row(i, cpu_board.cells, true)}    #{format_row(i, player_board.cells, false)}"
      },
      "\n"
    ]
  end
end
