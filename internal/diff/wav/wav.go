package wav

import (
	"context"
	"os"

	"github.com/emberfarkas/goctl/internal/diff/clap"
	"github.com/go-audio/wav"
	"github.com/go-bamboo/pkg/log"
	"github.com/spf13/cobra"
)

// Cmd represents the config command
var (
	Cmd = &cobra.Command{
		Use:   "wav",
		Short: "wav相关",
		Long:  `比较wav文件不同详情`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return diffWAV(cmd.Context())
		},
	}
)

func diffWAV(ctx context.Context) error {
	f1, err := os.Open(clap.Src)
	if err != nil {
		return err
	}
	d1 := wav.NewDecoder(f1)
	d1.ReadMetadata()
	log.Debugw("src", "NumChans", d1.NumChans, "BitDepth", d1.BitDepth, "SampleRate", d1.SampleRate, "WavAudioFormat", d1.WavAudioFormat, "AvgBytesPerSec", d1.AvgBytesPerSec)

	f2, err := os.Open(clap.Dst)
	if err != nil {
		return err
	}
	d2 := wav.NewDecoder(f2)
	d2.ReadMetadata()
	// log.Debugf("from: %v", d2.Metadata)
	log.Debugw("dst", "NumChans", d2.NumChans, "BitDepth", d2.BitDepth, "SampleRate", d2.SampleRate, "WavAudioFormat", d2.WavAudioFormat, "AvgBytesPerSec", d2.AvgBytesPerSec)

	return nil
}
