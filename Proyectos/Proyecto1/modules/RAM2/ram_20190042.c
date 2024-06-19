#include <linux/module.h>	
#include <linux/kernel.h>	
#include <linux/sched.h>

#include <linux/init.h>
#include <linux/proc_fs.h>
#include <asm/uaccess.h>
#include <linux/seq_file.h>
#include <linux/mm.h>

unsigned long rss;
const long kilobyte = 1024;

struct sysinfo si;

static void init_meminfo(void){
    si_meminfo(&si);
}

MODULE_LICENSE("GPL");
MODULE_AUTHOR("201900042");
MODULE_DESCRIPTION("Proyecto 1");

static int escribir_archivo(struct seq_file *archivo, void *v){
    init_meminfo();
    long total_ram = si.totalram*si.mem_unit;
    long ram_libre = si.freeram*si.mem_unit;
    long ram_en_uso = total_ram - ram_libre;
    long porcentaje_en_uso_cien = ram_en_uso *100;
    long porcentaje_en_uso = porcentaje_en_uso_cien/total_ram;
    seq_printf(archivo, "{\n\"Total_Ram\":%ld,\n", total_ram/kilobyte);
    seq_printf(archivo, "\"Ram_en_Uso\":%ld,\n", ram_en_uso/kilobyte);
    seq_printf(archivo, "\"Ram_Libre\":%ld,\n", ram_libre/kilobyte);
    seq_printf(archivo, "\"Porcentaje_en_uso\":%ld\n", porcentaje_en_uso);
    seq_printf(archivo, "}\n");

    return 0;
}

static int al_abrir(struct inode *inode, struct file *file){
    return single_open(file, escribir_archivo, NULL);
}

static struct proc_ops operaciones = 
{
    .proc_open = al_abrir,
    .proc_read = seq_read
};

static int __init modulo_insertar(void){
    proc_create("ram_201900042",0, NULL,&operaciones);
    printk(KERN_INFO "201900042\n");
    return 0;
}

static void __exit modulo_remove(void) {
    remove_proc_entry("ram_201900042", NULL);
    printk(KERN_INFO "Rodrigo Alejandro Hernández de León\n");
}

module_init(modulo_insertar);
module_exit(modulo_remove);