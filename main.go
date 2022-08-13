package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"test_smartm2m_muhammad_huzair/data"
	"test_smartm2m_muhammad_huzair/utils"
)

func main() {
	_, err := fmt.Fprintln(os.Stdout, utils.InitCmd)
	if err != nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
		}
		err = runCommand(cmdString)
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
		}
	}
}

func runCommand(commandStr string) error {
	var dataConfig data.ConfigFile

	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
		_, err := fmt.Fprintln(os.Stdout, "Are you sure you want to exit? (y/n)")
		if err != nil {
			return err
		}
		answer, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		answer = strings.TrimSuffix(answer, "\n")
		if answer != "y" {
			return nil
		}
		os.Exit(0)
	case "help":
		_, err := fmt.Fprintln(os.Stdout, utils.HelpString)
		if err != nil {
			return err
		}
		return nil
	case "generate":
		//id
		_, err := fmt.Fprintln(os.Stdout, "Choose fabricVersion (1.4.6 or 2.2.4) ?: ")
		if err != nil {
			return err
		}
		fabricVersion, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		fabricVersion = strings.TrimSuffix(fabricVersion, "\n")
		fabricVersion = strings.TrimSuffix(fabricVersion, "\r")

		//validation for fabricVersion (1.4.6 or 2.2.4)
		if !strings.Contains(fabricVersion, "1.4.6") && !strings.Contains(fabricVersion, "2.2.4") {
			return errors.New("invalid input for fabric version")
		}

		//enable or disable
		_, err = fmt.Fprintln(os.Stdout, "Enable or Disable Monitoring Log (y/n) ?: ")
		if err != nil {
			return err
		}
		enableDisableMonitoring, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		enableDisableMonitoring = strings.TrimSuffix(enableDisableMonitoring, "\n")
		enableDisableMonitoring = strings.TrimSuffix(enableDisableMonitoring, "\r")

		//add organization
	AddOrganization:
		_, err = fmt.Fprintln(os.Stdout, "Add Organization (y/n) ?: ")
		if err != nil {
			return err
		}
		answerAddOrganization, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		answerAddOrganization = strings.TrimSuffix(answerAddOrganization, "\n")
		answerAddOrganization = strings.TrimSuffix(answerAddOrganization, "\r")
		answerAddOrganization = strings.ToLower(answerAddOrganization)

		if answerAddOrganization == "y" {
			_, err = fmt.Fprintln(os.Stdout, "Name Channel ?: ")
			if err != nil {
				return err
			}
			organizationChannelName, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return err
			}
			organizationChannelName = strings.TrimSuffix(organizationChannelName, "\n")
			organizationChannelName = strings.TrimSuffix(organizationChannelName, "\r")

			_, err = fmt.Fprintln(os.Stdout, "Name Organization ?: ")
			if err != nil {
				return err
			}
			organizationName, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return err
			}
			organizationName = strings.TrimSuffix(organizationName, "\n")
			organizationName = strings.TrimSuffix(organizationName, "\r")

			_, err = fmt.Fprintln(os.Stdout, "Peers Organization (if more than 1, separate with comma): ")
			if err != nil {
				return err
			}
			organizationPeers, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return err
			}
			organizationPeers = strings.TrimSuffix(organizationPeers, "\n")
			organizationPeers = strings.TrimSuffix(organizationPeers, "\r")

			sanitizeStringWithComma := regexp.MustCompile(`^[a-zA-Z_\d ,]+$`)
			if !sanitizeStringWithComma.MatchString(organizationPeers) {
				return errors.New("invalid peers input")
			}

			listOrganizationArray := strings.Split(organizationPeers, ",")
			var orgs []data.Organization
			orgs = append(orgs, data.Organization{
				Name:  organizationName,
				Peers: listOrganizationArray,
			})
			channel := data.Channels{
				Name: organizationChannelName,
				Orgs: orgs,
			}

			dataConfig.Channels = append(dataConfig.Channels, channel)
			goto AddOrganization
		}

		//add channel codes
	AddChainCodes:
		_, err = fmt.Fprintln(os.Stdout, "Add ChainCodes (y/n) ?: ")
		if err != nil {
			return err
		}
		answerChainCodes, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		answerChainCodes = strings.TrimSuffix(answerChainCodes, "\n")
		answerChainCodes = strings.TrimSuffix(answerChainCodes, "\r")
		answerChainCodes = strings.ToLower(answerChainCodes)
		if answerChainCodes == "y" {
			_, err = fmt.Fprintln(os.Stdout, "Name ChainCode ?: ")
			if err != nil {
				return err
			}
			chainCodeName, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return err
			}
			chainCodeName = strings.TrimSuffix(chainCodeName, "\n")
			chainCodeName = strings.TrimSuffix(chainCodeName, "\r")

			_, err = fmt.Fprintln(os.Stdout, "Version ChainCode ?: ")
			if err != nil {
				return err
			}
			chainCodeVersion, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return err
			}
			chainCodeVersion = strings.TrimSuffix(chainCodeVersion, "\n")
			chainCodeVersion = strings.TrimSuffix(chainCodeVersion, "\r")

			_, err = fmt.Fprintln(os.Stdout, "Lang ChainCode ?: ")
			if err != nil {
				return err
			}
			chainCodeLang, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return err
			}
			chainCodeLang = strings.TrimSuffix(chainCodeLang, "\n")
			chainCodeLang = strings.TrimSuffix(chainCodeLang, "\r")

			_, err = fmt.Fprintln(os.Stdout, "Channel ChainCode ?: ")
			if err != nil {
				return err
			}
			chainCodeChannel, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return err
			}
			chainCodeChannel = strings.TrimSuffix(chainCodeChannel, "\n")
			chainCodeChannel = strings.TrimSuffix(chainCodeChannel, "\r")

			_, err = fmt.Fprintln(os.Stdout, "init ChainCode ?: ")
			if err != nil {
				return err
			}
			chainCodeInit, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return err
			}
			chainCodeInit = strings.TrimSuffix(chainCodeInit, "\n")
			chainCodeInit = strings.TrimSuffix(chainCodeInit, "\r")

			_, err = fmt.Fprintln(os.Stdout, "endorsement ChainCode ?: ")
			if err != nil {
				return err
			}
			chainCodeEndorsement, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return err
			}
			chainCodeEndorsement = strings.TrimSuffix(chainCodeEndorsement, "\n")
			chainCodeEndorsement = strings.TrimSuffix(chainCodeEndorsement, "\r")

			_, err = fmt.Fprintln(os.Stdout, "directory ChainCode ?: ")
			if err != nil {
				return err
			}
			chainCodeDirectory, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				return err
			}
			chainCodeDirectory = strings.TrimSuffix(chainCodeDirectory, "\n")
			chainCodeDirectory = strings.TrimSuffix(chainCodeDirectory, "\r")

			chainCode := data.ChainCodes{
				Name:        chainCodeName,
				Version:     chainCodeVersion,
				Lang:        chainCodeLang,
				Channel:     chainCodeChannel,
				Init:        chainCodeInit,
				Endorsement: chainCodeEndorsement,
				Directory:   chainCodeDirectory,
			}
			dataConfig.ChainCodes = append(dataConfig.ChainCodes, chainCode)

			goto AddChainCodes
		}

		//collect data
		if enableDisableMonitoring == "y" {
			enableDisableMonitoring = "enabled"
		} else if enableDisableMonitoring == "n" {
			enableDisableMonitoring = "disabled"
		}

		dataConfig.Global.FabricVersion = fabricVersion
		dataConfig.Global.Monitoring.Loglevel = enableDisableMonitoring

		err = data.GenerateConfigFile(dataConfig)

		_, err = fmt.Fprintln(os.Stdout)
		if err != nil {
			return err
		}
		return nil
	default:
		_, err := fmt.Fprintf(os.Stdout, "%s is unknown command\n", arrCommandStr[0])
		_, err = fmt.Fprintf(os.Stdout, "this is the list command you can use \n")
		_, err = fmt.Fprintln(os.Stdout, utils.HelpString)
		_, err = fmt.Fprintf(os.Stdout, "you can list the command of the quizmaster by typing help \n")
		if err != nil {
			return err
		}
	}
	command := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	return command.Run()
}
