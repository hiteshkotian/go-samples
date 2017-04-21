package script

import "os/exec"

func ExecuteCommand (command string, arg ...string) (string, error) {
   cmd := exec.Command(command, arg...)

   result_byte, err := cmd.Output()

   if err != nil {
      return "", err
   }

   str_output := string (result_byte[:])
   return str_output, nil
   //return string (result_byte[:]), nil
}
