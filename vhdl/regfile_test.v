/* time scale */
`timescale 1ns/1ps

/* headerfile */
`include "regfile.h"

/* module */
module regfile_test;
    /* io define */
    reg            clk;
    reg            reset_;
    reg [`AddrBus] addr;
    reg [`DataBus] d_in;
    reg            we_;
    wire [`DataBus] d_out;
    /* internal var */
    integer        i;
    /* def simu cycles */
    parameter      STEP = 100.0000; //10M
    /* instanc timer */
    always #(STEP / 2) begin
        clk <= ~clk;
    end

    /* instanc test obj*/
    regfile regfile (
        /* timer and reset */
        .clk (clk),
        .reset_ (reset_),
        /* access port */
        .addr (addr),
        .d_in (d_in),
        .we_ (we_),
        .d_out (d_out)
    );

    /* test case */
    initial begin
        # 0 begin
            clk <= `HIGH;
            reset_ <= `ENABLE_;
            addr <= {`ADDR_W{1'b0}};
            d_in <= {`DATA_W{1'b0}};
            we_ <= `DISABLE_;
        end
        # (STEP * 3 / 4)
        # STEP begin
            reset_ <= `DISABLE_;
        end
        # STEP begin
            for (i=0; i<`DATA_D; i=i+1) begin
                # STEP begin
                    addr <= i;
                    d_in <= i;
                    we_ <= `ENABLE_;
                end
                # STEP begin
                    addr <= {`ADDR_W{1'b0}};
                    d_in <= {`DATA_W{1'b0}};
                    we_ <= `DISABLE_;
                    if (d_out == i) begin
                        $display($time, " ff[%d] Read/Write Check OK !", i);
                    end else begin
                        $display($time, " ff[%d] Read/Write Check NG !", i);
                    end
                end
            end
        # STEP begin
            $finish;
        end
    end

    /* wave output */
    // initial begin
    //     // $dumpfile("C:\\regfile.vcd");
    //     // $dumpvars(0, regfile);
    // end

endmodule