(require :sb-bsd-sockets)
(use-package :sb-bsd-sockets)

(defun get-hostname ()
  (machine-instance))

(defun linux-p ()
  (find :linux *features*))

(defun windows-p ()
  (find :win32 *features*))

(defun macos-p ()
  (find :darwin *features*))

(defun run-plugin (code)
  (eval (read-from-string code)))

(defun download-plugin (url)
  )

(defun heartbeat ()
  )

;; Remove the test function in production
(defun test ()
  (run-plugin "(sb-ext:run-program \"/usr/bin/whoami\" () :output t)"))
