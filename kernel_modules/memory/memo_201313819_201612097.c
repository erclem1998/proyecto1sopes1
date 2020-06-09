#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <asm/uaccess.h>
#include <linux/hugetlb.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/fs.h>

#define BUFSIZE     150

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Información de memoria");
MODULE_AUTHOR("Erick Lemus - 201612097\nJavier Solares - 201313819");

struct sysinfo inf;

static int write_file(struct seq_file * archivo, void *v){
    si_meminfo(&inf);
    long memoriatotal=(inf.totalram*4);
    long memorialibre=(inf.freeram*4);
    seq_printf(archivo,"\n");
    seq_printf(archivo,"LABORATORIO SISTEMAS OPERATIVOS 1\n");
    seq_printf(archivo,"JUNIO 2020\n");
    seq_printf(archivo,"CESAR JAVIER SOLARES OROZCO - 201313819\n");
    seq_printf(archivo,"ERICK ALEXANDER LEMUS MORALES - 201612097\n");
    seq_printf(archivo,"PROYECTO 1 - MODULO DE MEMORIA\n");
    seq_printf(archivo,"\n");
    seq_printf(archivo,"TOTAL MEMORIA -> \t  %8lu kb - %8lu mb \n", memoriatotal, memoriatotal/1024);
    seq_printf(archivo,"MEMORIA LIBRE -> \t %8lu KB - %8lu MB \n", memorialibre, memorialibre/1024);
    seq_printf(archivo,"MEMORIA EN USO -> \t %i %%\n", (memorialibre * 100)/memoriatotal);
    seq_printf(archivo,"\n");
    return 0;
}

static int abrir(struct inode  *inode, struct file *file){
    return single_open(file,write_file, NULL);
}

static struct file_operations ops =
{
    .open=abrir,
    .read=seq_read
    /* data */
};

static int iniciar(void){
    proc_create("memo_201313819_201612097",0, NULL,&ops);
    printk("\nCarnet1: 201313819, Carnet2: 201612097\n");
    return 0;
}

static void salir(void){
    remove_proc_entry("memo_201313819_201612097",NULL);
    printk("\nSistemas Operativos 1\n");
}

module_init(iniciar);
module_exit(salir);

