package main

import (
	"testing"
)

func TestMoveDisk(t *testing.T) {
	tests := []struct {
		name           string
		initialTrails  []Trail
		from           int
		to             int
		expectedTrails []Trail
		expectError    bool
	}{
		{
			name: "Move disk successfully",
			initialTrails: []Trail{
				{
					Name: "1",
					Disks: []Disk{
						{Name: "Disk1", Color: "#FFFFFF", Size: 1},
					},
				},
				{
					Name:  "2",
					Disks: []Disk{},
				},
			},
			from: 1,
			to:   2,
			expectedTrails: []Trail{
				{
					Name:  "1",
					Disks: []Disk{},
				},
				{
					Name: "2",
					Disks: []Disk{
						{Name: "Disk1", Color: "#FFFFFF", Size: 1},
					},
				},
			},
			expectError: false,
		},
		{
			name: "No disk to move",
			initialTrails: []Trail{
				{
					Name:  "1",
					Disks: []Disk{},
				},
				{
					Name:  "2",
					Disks: []Disk{},
				},
			},
			from: 1,
			to:   2,
			expectedTrails: []Trail{
				{
					Name:  "1",
					Disks: []Disk{},
				},
				{
					Name:  "2",
					Disks: []Disk{},
				},
			},
			expectError: true,
		},
		{
			name: "Cannot move larger disk onto smaller disk",
			initialTrails: []Trail{
				{
					Name: "1",
					Disks: []Disk{
						{Name: "Disk1", Color: "#FFFFFF", Size: 2},
					},
				},
				{
					Name: "2",
					Disks: []Disk{
						{Name: "Disk2", Color: "#FFFFFF", Size: 1},
					},
				},
			},
			from: 1,
			to:   2,
			expectedTrails: []Trail{
				{
					Name: "1",
					Disks: []Disk{
						{Name: "Disk1", Color: "#FFFFFF", Size: 2},
					},
				},
				{
					Name: "2",
					Disks: []Disk{
						{Name: "Disk2", Color: "#FFFFFF", Size: 1},
					},
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trailManager := &TrailManager{
				Trails: tt.initialTrails,
			}

			trailManager.MoveDisk(tt.from, tt.to)

			for i, trail := range trailManager.Trails {
				if len(trail.Disks) != len(tt.expectedTrails[i].Disks) {
					t.Errorf("expected %d disks in trail %d, got %d", len(tt.expectedTrails[i].Disks), i+1, len(trail.Disks))
				}
				for j, disk := range trail.Disks {
					if disk != tt.expectedTrails[i].Disks[j] {
						t.Errorf("expected disk %v in trail %d, got %v", tt.expectedTrails[i].Disks[j], i+1, disk)
					}
				}
			}
		})
	}
}

func TestHexTo256Color(t *testing.T) {
	tests := []struct {
		hexColor string
		expected int
	}{
		{"#000000", 16},  // Black
		{"#FFFFFF", 231}, // White
		{"#FF0000", 196}, // Red
		{"#00FF00", 46},  // Green
		{"#0000FF", 21},  // Blue
		{"#FFFF00", 226}, // Yellow
		{"#FFA500", 214}, // Orange
	}

	for _, test := range tests {
		result := hexTo256Color(test.hexColor)
		if result != test.expected {
			t.Errorf("hexTo256Color(%s) = %d; want %d", test.hexColor, result, test.expected)
		}
	}
}
