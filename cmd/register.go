package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "todo를 등록할 수 있는 커맨드",
	Long: `todo를 등록할 수 있습니다.
플래그를 통해 todo의 내용을 지정할 수 있습니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("title", "t", "", "todo의 제목의 값을 넣는 플래그입니다.")
	registerCmd.Flags().StringP("details", "d", "", "todo의 세부 내용을 넣는 플래그입니다.")
	registerCmd.Flags().StringP("startDate", "", "", "todo의 시작일자를 넣는 플래그입니다.")
	registerCmd.Flags().StringP("endDate", "", "", "todo의 종료일자를 넣는 플래그입니다.")
	registerCmd.Flags().StringP("priority", "p", "", "todo의 우선 순위를 넣는 플래그입니다.")
}
