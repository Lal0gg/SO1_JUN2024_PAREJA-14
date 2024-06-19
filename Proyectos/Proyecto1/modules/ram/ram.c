#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/mm.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/version.h>
#include <linux/vmstat.h>
#include <linux/mmzone.h>

struct sysinfo si;

static void init_meminfo(void) {
    si_meminfo(&si);
}

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de RAM, Laboratorio Sistemas Operativos 1");
MODULE_AUTHOR("Lal0gg");

static unsigned long calculate_mem_available(void)
{
    long cached, available;
    unsigned long pagecache, reclaimable;

    cached = global_node_page_state(NR_FILE_PAGES) -
             global_node_page_state(NR_SHMEM);

    pagecache = global_node_page_state(NR_ACTIVE_FILE) +
                global_node_page_state(NR_INACTIVE_FILE);
    reclaimable = global_node_page_state(NR_SLAB_RECLAIMABLE_B) +
                  global_node_page_state(NR_KERNEL_MISC_RECLAIMABLE);

    available = si.freeram + cached + reclaimable -
                min(pagecache / 2, si.totalram - si.freeram - reclaimable);

    return available * si.mem_unit / 1024;
}

static int escribir_archivo(struct seq_file *archivo, void *v)
{
    init_meminfo();

    // Calcular la RAM total en KB
    unsigned long total_ram = (si.totalram * si.mem_unit) / 1024;

    // Calcular la RAM libre en KB
    unsigned long free_ram = (si.freeram * si.mem_unit) / 1024;

    // Calcular la RAM disponible en KB
    unsigned long available_ram = calculate_mem_available();

    // Calcular porcentaje de RAM libre
    unsigned long percent_free = (available_ram * 100) / total_ram;

    // Calcular porcentaje de RAM en uso
    unsigned long percent_used = 100 - percent_free;

    // Escribir los valores en el archivo
    seq_printf(archivo, "RAMfree,%lu\n", free_ram);
    seq_printf(archivo, "RamUsed,%lu\n", percent_used);

    return 0;
}

static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

#if LINUX_VERSION_CODE >= KERNEL_VERSION(5,6,0)
static const struct proc_ops operaciones =
{
   .proc_open = al_abrir,
   .proc_read = seq_read,
   .proc_lseek = seq_lseek,
   .proc_release = single_release
};
#else
static const struct file_operations operaciones =
{
   .open = al_abrir,
   .read = seq_read,
   .llseek = seq_lseek,
   .release = single_release
};
#endif

static int __init insertar(void)
{
    proc_create("ram_so1_jun2024", 0, NULL, &operaciones);
    printk(KERN_INFO "201900647\n");
    return 0;
}

static void __exit remover(void)
{
    remove_proc_entry("ram_so1_jun2024", NULL);
    printk(KERN_INFO "Laboratorio Sistemas Operativos 1\n");
}

module_init(insertar);
module_exit(remover);
