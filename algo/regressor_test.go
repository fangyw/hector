package algo

import (
    "testing"
)

func TestRegressorOnSin(t *testing.T) {
    algos := []string{"gp"}

    params := make(map[string]string)
    params["dim"] = "1"

    for _, algo := range algos {
        train_dataset := SinusoidalDataSet(100)
        test_dataset := SinusoidalDataSet(50)
        regressor := GetRegressor(algo)
        regressor.Init(params)
        rmse, _ := RegAlgorithmRunOnDataSet(regressor, train_dataset, test_dataset, "", params)

        t.Logf("rmse of %s in sinusoidal dataset is %f", algo, rmse)
        if rmse > 0.1 {
            t.Error("rmse less than 0.1 in sinusoidal dataset")
        }
    }
}
