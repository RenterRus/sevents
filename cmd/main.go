package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "sevents",
	Short: "Root command",
	Long:  "Available flags: -m (mongodb address) -d (database name) -c (collection name) -l (http server listening address)",
	Run: RunSevents,
}

//Хотелось бы тут использовать считывание конфиг файла, но, не уверен, что такой оверинжениринг уместен, так что,
//просто, кобра с флагами
func init(){
	rootCmd.Flags().StringP("listen", "l", "127.0.0.1:9099", "Listen HTTP")
	rootCmd.Flags().StringP("mongo", "m", "mongodb://127.0.0.1", "Address MongoDB")
	rootCmd.Flags().StringP("db", "d", "storing", "DB name")
	rootCmd.Flags().StringP("collection", "c", "events", "Collection name")
	rootCmd.Execute()
}
func main(){
}