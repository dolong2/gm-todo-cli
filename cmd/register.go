package cmd

import (
	"github.com/spf13/cobra"
	"gm-todo/model"
	"strings"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "todo를 등록할 수 있는 커맨드",
	Long: `todo를 등록할 수 있습니다.
플래그를 통해 todo의 내용을 지정할 수 있습니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workToDo, err := cmd.Flags().GetString("title")
		if err != nil {
			return err
		}

		details, err := cmd.Flags().GetString("details")
		if err != nil {
			return err
		}

		startDate, err := cmd.Flags().GetString("startDate")
		if err != nil {
			return err
		}

		endDate, err := cmd.Flags().GetString("endDate")
		if err != nil {
			return err
		}

		priority, err := cmd.Flags().GetString("priority")
		if err != nil {
			return err
		}

		command := strings.Join([]string{workToDo, details, startDate, endDate, priority}, ", ")

		model.SendRequest(command)
		return nil
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
