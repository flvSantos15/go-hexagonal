Run: func(cmd *cobra.Command, args []string) {
	res, err := cli.Run(&productService, action, productId, productName, productPrice)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}

init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", 0)
	cliCmd.Flags().StringVarP(&productId, "id", "i", "enable", 0)
	cliCmd.Flags().StringVarP(&productName, "product", "n", "enable", 0)
	cliCmd.Flags().StringVarP(&productPrice, "price", "p", "enable", 0)
}