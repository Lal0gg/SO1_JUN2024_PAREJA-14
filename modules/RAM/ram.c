// Libreria para interactuar con modulos de kernel
#include <linux/module.h>
// Librería para obtener información de kernel
#include <linux/kernel.h>
// Información de la memoria RAM
#include <linux/mm.h>

// Header para los macros module_init y module_exit
#include <linux/init.h>
// Header necesario porque se usara proc_fs
#include <linux/proc_fs.h>
/* for copy_from_user */
#include <asm/uaccess.h>
/* Header para usar la lib seq_file y manejar el archivo en /proc*/
#include <linux/seq_file.h>
// Incluir version.h para obtener la version del kernel
#include <linux/version.h>

// Obtener las estadísticas del sistema
struct sysinfo si;

static void init_meminfo(void) {
    si_meminfo(&si);
}

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de RAM, Laboratorio Sistemas Operativos 1");
MODULE_AUTHOR("Lal0gg");

// Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int escribir_archivo(struct seq_file *archivo, void *v)
{
    init_meminfo();

    // Calcular la RAM utilizada y la RAM libre en KB
    unsigned long total_ram = si.totalram * si.mem_unit / 1024;
    unsigned long free_ram = si.freeram * si.mem_unit / 1024;
    unsigned long used_ram = total_ram - free_ram;

    // Escribir los valores en el archivo
    seq_printf(archivo, "RAM utilizada: %lu kB\n", used_ram);
    seq_printf(archivo, "RAM libre: %lu kB\n", free_ram);

    return 0;
}

// Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

// Si el kernel es 5.6 o mayor se usa la estructura proc_ops
#if LINUX_VERSION_CODE >= KERNEL_VERSION(5,6,0)
static struct proc_ops operaciones =
{
    .proc_open = al_abrir,
    .proc_read = seq_read,
    .proc_lseek = seq_lseek,
    .proc_release = single_release
};
#else
static struct file_operations operaciones =
{
    .open = al_abrir,
    .read = seq_read,
    .llseek = seq_lseek,
    .release = single_release
};
#endif

// Funcion a ejecutar al insertar el modulo en el kernel con insmod
static int __init insertar(void)
{
    proc_create("ram_201900647", 0, NULL, &operaciones);
    printk(KERN_INFO "201900647\n");
    return 0;
}

// Funcion a ejecutar al remover el modulo del kernel con rmmod
static void __exit remover(void)
{
    remove_proc_entry("ram_201900647", NULL);
    printk(KERN_INFO "Laboratorio Sistemas Operativos 1\n");
}

module_init(insertar);
module_exit(remover);