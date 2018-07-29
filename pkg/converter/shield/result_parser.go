package shield

import (
	"strings"

	"github.com/yaptide/yaptide/pkg/converter"
)

// ParseResults will parse results of shield simulation.
func ParseResults(
	files map[string]string, simulationContext *SerializationContext,
) (*converter.Result, error) {
	log.Info("[Parser][Results] Start shield parser.")

	simulationResult := converter.NewEmptyResult()

	for bdoFile, content := range files {
		if strings.Contains(bdoFile, ".bdo") {
			log.Debug("[Parser][Results] Start parsing result file %s", bdoFile)
			parser := newBdoParser(bdoFile[:len(bdoFile)-4], []byte(content), *simulationContext)
			parseErr := parser.Parse()
			if parseErr != nil {
				log.Warning("[Parser][Results] file parsing error %s", parseErr.Error())
			}
			simulationResult.AddDetectorResults(parser.Results)
		}
	}

	log.Info("[Parser][Results] Finished shield parser")
	return &simulationResult, nil
}
