package main

func main() {
	rootCmd := NewRootCmd()
	rootCmd.AddCommand(NewProxyCmd())
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
