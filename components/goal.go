package components

import (
	"copatoon/utils"
	"log/slog"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Goal(width, height int) [][][]byte {
	goalHeight := height / 3
	if (height-goalHeight)%2 == 1 {
		goalHeight--
	}
	halfLeftHeight := (height - goalHeight) / 2
	utils.Assert(goalHeight+2*halfLeftHeight == height,
		"GAME",
		"The sum of heights is not equal to the total height",
		slog.Int("halfLeftHeight", halfLeftHeight),
		slog.Int("goalHeight", goalHeight),
		slog.Int("totalHeight", height),
	)

	goal := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		Padding(0, 1).
		Margin(halfLeftHeight-1, width-4, halfLeftHeight-1, 0).
		Render(strings.Repeat("\n", goalHeight-1))

	utils.Assert(strings.Count(goal, "\n") == height-1,
		"GAME",
		"The goal view height is not equal to the total height",
		slog.Int("goalHeight", strings.Count(goal, "\n")),
		slog.Int("totalHeight", height),
	)

	utils.Assert(len(strings.Split(goal, "\n")[0]) == width,
		"GAME",
		"The goal view width is not equal to the total width",
		slog.Int("goalWidth", len(strings.Split(goal, "\n")[0])),
		slog.Int("totalWidth", width),
	)

	// TODO: Make non-space be stored as 3-byte slices, so that I can easily identify when rendering
	goalPartsBytes := [][][]byte{}
	goalParts := strings.Split(goal, "\n")
	for i := range height {
		j := 0
		goalPartsBytes = append(goalPartsBytes, make([][]byte, 0))
		for j < len(goalParts[i]) {
			if ' ' == goalParts[i][j] {
				goalPartsBytes[i] = append(goalPartsBytes[i], []byte{' '})
				j++
				continue
			}

			goalPartsBytes[i] = append(goalPartsBytes[i], []byte{
				goalParts[i][j],
				goalParts[i][j+1],
				goalParts[i][j+2],
			})
			j += 3
		}
	}

	utils.Assert(len(goalPartsBytes) == height,
		"GAME",
		"The goal view width is not equal to the total width",
		slog.Int("goalHeight", len(goalPartsBytes)),
		slog.Int("totalHeight", height),
	)

	utils.Assert(len(goalPartsBytes[0]) == width,
		"GAME",
		"The goal view width is not equal to the total width",
		slog.Int("goalWidth", len(goalPartsBytes[0])),
		slog.Int("totalWidth", width),
	)

	return goalPartsBytes
}
