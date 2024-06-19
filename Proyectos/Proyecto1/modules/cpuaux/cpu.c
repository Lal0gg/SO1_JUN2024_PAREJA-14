#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/sched.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <asm/uaccess.h>
#include <linux/seq_file.h>
#include <linux/mm.h>

struct task_struct *cpu;
struct list_head *lstProcess;
struct task_struct *child;
unsigned long rss;

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de CPU para el Lab de Sopes 1");
MODULE_AUTHOR("Dani :)");

static int escribir_archivo(struct seq_file *archivo, void *v) {
    seq_printf(archivo, "{\n");

    seq_printf(archivo, "\"processes\": [\n");

    bool first = true;

    for_each_process(cpu) {
        if (!first)
            seq_printf(archivo, ",\n");
        else
            first = false;

        seq_printf(archivo, "{\n");
        seq_printf(archivo, "\"pid\": %d,\n", cpu->pid);
        seq_printf(archivo, "\"name\": \"%s\",\n", cpu->comm);
        seq_printf(archivo, "\"state\": %lu,\n", cpu->__state);

        if (cpu->mm) {
            rss = get_mm_rss(cpu->mm) << PAGE_SHIFT;
            seq_printf(archivo, "\"ram\": %lu,\n", rss);
        } else {
            seq_printf(archivo, "\"ram\": \"\",\n");
        }

        seq_printf(archivo, "\"user\": %d,\n", cpu->cred->user->uid);

        seq_printf(archivo, "\"children\": [\n");
        bool child_first = true;
        list_for_each(lstProcess, &(cpu->children)) {
            if (!child_first)
                seq_printf(archivo, ",\n");
            else
                child_first = false;

            child = list_entry(lstProcess, struct task_struct, sibling);
            seq_printf(archivo, "{\n");
            seq_printf(archivo, "\"pid\": %d,\n", child->pid);
            seq_printf(archivo, "\"name\": \"%s\",\n", child->comm);
            seq_printf(archivo, "\"state\": %lu,\n", child->__state);

            if (child->mm) {
                rss = get_mm_rss(child->mm) << PAGE_SHIFT;
                seq_printf(archivo, "\"ram\": %lu,\n", rss);
            } else {
                seq_printf(archivo, "\"ram\": \"\",\n");
            }

            seq_printf(archivo, "\"user\": %d\n", child->cred->user->uid);
            seq_printf(archivo, "}\n");
        }
        seq_printf(archivo, "]\n");

        seq_printf(archivo, "}\n");
    }

    seq_printf(archivo, "]\n");
    seq_printf(archivo, "}\n");

    return 0;
}

static int al_abrir(struct inode *inode, struct file *file) {
    return single_open(file, escribir_archivo, NULL);
}

static struct proc_ops operaciones =
{
    .proc_open = al_abrir,
    .proc_read = seq_read
};

static int _insert(void) {
    proc_create("cpu_201800722", 0, NULL, &operaciones);
    printk(KERN_INFO "Jose Daniel Velasquez Orozco\n");
    return 0;
}

static void _remove(void) {
    remove_proc_entry("cpu_201800722", NULL);
    printk(KERN_INFO "Segundo Semestre 2023\n");
}

module_init(_insert);
module_exit(_remove);
