#include <linux/module.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/sched/signal.h>
#include <linux/seq_file.h>
#include <linux/fs.h>
#include <linux/sched.h>
#include <linux/mm.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Eduardo Gonzalez");
MODULE_DESCRIPTION("Informacion cpu");
MODULE_VERSION("1.0");

struct task_struct *task;       // sched.h para tareas/procesos
struct task_struct *task_child; // index de tareas secundarias
struct list_head *list;         // lista de cada tareas

static const char* state_to_string(char state) {
    switch (state) {
        case 'R':
            return "Running";
        case 'S':
            return "Sleeping";
        case 'D':
            return "Sleeping";
        case 'T':
            return "Stopped";
        case 't':
            return "Traced";
        case 'Z':
            return "Zombie";
        case 'X':
            return "Dead";
        default:
            return "Unknown";
    }
}

static int escribir_a_proc(struct seq_file *file_proc, void *v) {
    int running = 0;
    int sleeping = 0;
    int zombie = 0;
    int stopped = 0;
    unsigned long rss;
    unsigned long total_ram_pages;

    total_ram_pages = totalram_pages();
    if (!total_ram_pages) {
        pr_err("No memory available\n");
        return -EINVAL;
    }

    unsigned long total_cpu_time = 0;

    for_each_process(task) {
        total_cpu_time += task->utime + task->stime;
    }

    //---------------------------------------------------------------------------
    seq_printf(file_proc, "{\n\"CpuUsed\":%u,\n", jiffies_to_msecs(get_jiffies_64()));
    seq_printf(file_proc, "\"CpuPercent\":%u,\n", (total_cpu_time * 100) / jiffies_to_msecs(get_jiffies_64()));
    seq_printf(file_proc, "\"processes\":[\n");
    int b = 0;

    for_each_process(task) {
        if (task->mm) {
            rss = get_mm_rss(task->mm) << PAGE_SHIFT;
        } else {
            rss = 0;
        }
        if (b == 0) {
            seq_printf(file_proc, "{");
            b = 1;
        } else {
            seq_printf(file_proc, ",{");
        }
        seq_printf(file_proc, "\"pid\":%d,\n", task->pid);
        seq_printf(file_proc, "\"name\":\"%s\",\n", task->comm);
        seq_printf(file_proc, "\"user\": %u,\n", from_kuid(&init_user_ns, task->cred->uid));
        seq_printf(file_proc, "\"state\":\"%s\",\n", state_to_string(task_state_to_char(task)));
        int porcentaje = (rss * 100) / total_ram_pages;
        seq_printf(file_proc, "\"ram\":%d,\n", porcentaje);

        seq_printf(file_proc, "\"child\":[\n");
        int a = 0;
        list_for_each(list, &(task->children)) {
            task_child = list_entry(list, struct task_struct, sibling);
            if (a != 0) {
                seq_printf(file_proc, ",{");
                seq_printf(file_proc, "\"pid\":%d,\n", task_child->pid);
                seq_printf(file_proc, "\"name\":\"%s\",\n", task_child->comm);
                seq_printf(file_proc, "\"state\":\"%s\",\n", state_to_string(task_state_to_char(task_child)));
                seq_printf(file_proc, "\"pidPadre\":%d\n", task->pid);
                seq_printf(file_proc, "}\n");
            } else {
                seq_printf(file_proc, "{");
                seq_printf(file_proc, "\"pid\":%d,\n", task_child->pid);
                seq_printf(file_proc, "\"name\":\"%s\",\n", task_child->comm);
                seq_printf(file_proc, "\"state\":\"%s\",\n", state_to_string(task_state_to_char(task_child)));
                seq_printf(file_proc, "\"pidPadre\":%d\n", task->pid);
                seq_printf(file_proc, "}\n");
                a = 1;
            }
        }
        a = 0;
        seq_printf(file_proc, "\n]");

        switch (task_state_to_char(task)) {
            case 'R':
                running++;
                break;
            case 'S':
            case 'D':
                sleeping++;
                break;
            case 'Z':
                zombie++;
                break;
            case 'T':
            case 't':
                stopped++;
                break;
        }
        seq_printf(file_proc, "}\n");
    }
    b = 0;
    seq_printf(file_proc, "],\n");
    seq_printf(file_proc, "\"running\":%d,\n", running);
    seq_printf(file_proc, "\"sleeping\":%d,\n", sleeping);
    seq_printf(file_proc, "\"zombie\":%d,\n", zombie);
    seq_printf(file_proc, "\"stopped\":%d,\n", stopped);
    seq_printf(file_proc, "\"total\":%d\n", running + sleeping + zombie + stopped);
    seq_printf(file_proc, "}\n");
    return 0;
}

static int abrir_aproc(struct inode *inode, struct file *file) {
    return single_open(file, escribir_a_proc, NULL);
}

static const struct proc_ops archivo_operaciones = {
    .proc_open = abrir_aproc,
    .proc_read = seq_read
};

static int __init modulo_init(void) {
    proc_create("cpu_so1_1s2024", 0, NULL, &archivo_operaciones);
    printk(KERN_INFO "Insertar Modulo CPU\n");
    return 0;
}

static void __exit modulo_cleanup(void) {
    remove_proc_entry("cpu_so1_1s2024", NULL);
    printk(KERN_INFO "Remover Modulo CPU\n");
}

module_init(modulo_init);
module_exit(modulo_cleanup);