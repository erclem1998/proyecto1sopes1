/*#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/module.h>
#include <linux/list.h>
#include <linux/types.h>
#include <linux/slab.h>
#include <linux/sched.h>
#include <linux/string.h>
#include <linux/fs.h>
#include <linux/seq_file.h>
#include <linux/proc_fs.h>
#include <linux/mm.h>
#include <linux/sched/signal.h>
*/
/*void recorrerHijos(struct seq_file *sf, struct task_struct *ts, int contador)
{
    struct list_head *list;
    struct task_struct *task;
    
    list_for_each(list, &ts->children)
    {
        task = list_entry(list, struct task_struct, sibling);
        pr_info("Nombre: %s PID: %d Estado: %d\n", ts->comm, ts->pid, ts->state);
        recorrerHijos(sf, task, contador + 1);
    }
}*/
/*

typedef struct procesos
{
    int pid;
    char nombre[150];
    int estado;
    int pidpadre;
};

void procs_info_print(void)
{
    struct task_struct *task_list;
    size_t process_counter = 0;
    int tam=0;
    for_each_process(task_list)
    {
        pr_info("Nombre: %s \t PID: %d \t Estado: %d \t PID Padre: %d\n", task_list->comm, task_list->pid, task_list->state, task_list->real_parent->pid);
        ++process_counter;
        tam++;
    }
    printk(KERN_INFO "== Number of process: %zu\n", process_counter);
}

int init_module(void)
{
    printk(KERN_INFO "[ INIT ==\n");*/
    /*struct task_struct *root = current;
    while (root->pid != 1)
    {
        root = root->parent;
    }
    struct seq_file *sf;
    recorrerHijos(sf, root, 1);
    */
/*    procs_info_print();

    return 0;
}

void cleanup_module(void)
{
    printk(KERN_INFO "== CLEANUP ]\n");
}

MODULE_LICENSE("MIT");
*/

#include <linux/kernel.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/sched/signal.h>
#include <linux/sched.h>

#include <linux/list.h>
#include <linux/types.h>
#include <linux/slab.h>
#include <linux/string.h>
#include <linux/fs.h>
#include <linux/seq_file.h>
#include <linux/proc_fs.h>
#include <linux/mm.h>
 
 
struct task_struct *task;        /*    Structure defined in sched.h for tasks/processes    */
struct task_struct *task_child;        /*    Structure needed to iterate through task children    */
struct list_head *list;            /*    Structure needed to iterate through the list in each task->children struct    */

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("CPU");
MODULE_AUTHOR("Javier Solares 201313819 --- Erick Lemus 201612097");

int iterate_init(struct seq_file * archivo)                    /*    Init Module    */
{
    //printk(KERN_INFO "%s","LOADING MODULE\n");    /*    good practice to log when loading/removing modules    */
    seq_printf(archivo,"---------OBTENIENDO PROCESOS---------\n");
    for_each_process( task ){            /*    for_each_process() MACRO for iterating through each task in the os located in linux\sched\signal.h    */
        seq_printf(archivo,"\nPARENT PID: %d PROCESS: %s STATE: %ld \n",task->pid, task->comm, task->state);/*    log parent id/executable name/state    */
        list_for_each(list, &task->children){                        /*    list_for_each MACRO to iterate through task->children    */
            task_child = list_entry( list, struct task_struct, sibling );    /*    using list_entry to declare all vars in task_child struct    */
     
            char estado=79; //otro estado
            if(task_child->state==TASK_RUNNING){
                estado=82;
            }
            else if(task_child->state==__TASK_STOPPED){
                estado="S";
            }
            else if(task_child->state==TASK_INTERRUPTIBLE){
                estado=73;
            }
            else if(task_child->state==TASK_UNINTERRUPTIBLE){
                estado="U";
            }
            else if(task_child->exit_state==EXIT_ZOMBIE){
                estado="Z";
            }
            else if(task_child->state==TASK_DEAD){
                estado="D";
            }
            seq_printf(archivo, "\nCHILD OF %s[%d] PID: %d PROCESS: %s ESTADO: %c \n",task->comm, task->pid, /*    log child of and child pid/name/state    */
                task_child->pid, task_child->comm, estado);
        }
        seq_printf(archivo,"-----------------------------------------------------\n");    /*for aesthetics*/
    }    
     
 
    return 0;
 
}

static int write_file(struct seq_file * archivo, void *v){
    seq_printf(archivo,"\n");
    seq_printf(archivo,"      -------------------------------------\n");
    seq_printf(archivo,"      |     PROYECTO 1 - MODULO DE CPU    |\n");
    seq_printf(archivo,"      | LABORATORIO SISTEMAS OPERATIVOS 1 |\n");
    seq_printf(archivo,"      |           JUNIO 2020              |\n");
    seq_printf(archivo,"      -------------------------------------\n");
    seq_printf(archivo,"\n");
    seq_printf(archivo,"      CESAR JAVIER SOLARES OROZCO - 201313819\n");
    seq_printf(archivo,"     ERICK ALEXANDER LEMUS MORALES - 201612097\n");
    seq_printf(archivo,"\n");
    return iterate_init(archivo);
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
    proc_create("cpu_201313819_201612097",0, NULL,&ops);
    printk(KERN_INFO "\nCarnet1: 201313819, Carnet2: 201612097\n");
    return 0;
}

static void salir(void){
    remove_proc_entry("cpu_201313819_201612097",NULL);
    printk(KERN_INFO "\nSistemas Operativos 1\n");
}
                /*    End of Init Module    */
     
void cleanup_exit(void)        /*    Exit Module    */
{
 
 
    printk(KERN_INFO "%s","REMOVING MODULE\n");
 
}                /*    End of Exit Module    */
 
module_init(iniciar);    /*    Load Module MACRO    */
module_exit(salir);    /*    Remove Module MACRO    */