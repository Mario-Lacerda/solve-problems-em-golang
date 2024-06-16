func desvioPadrao(numeros []int) float64 {
    media := media(numeros)
    variancia := 0.0
    for _, numero := range numeros {
        variancia += math.Pow(float64(numero)-media, 2)
    }
    variancia /= float64(len(numeros))
    return math.Sqrt(variancia)
}

