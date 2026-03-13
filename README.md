**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-13 21:32:40 UTC（2026-03-14 05:32:40 UTC+8）

**代理总数：212**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 211 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 212 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 671ms | ✓ 903ms | ✓ 741ms | ✓ 1339ms | ✓ 906ms | http |
| 5.129.206.247:8888 | ✓ 1356ms | 否 | ✓ 1701ms | 否 | ✓ 1821ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1590ms | ✓ 1538ms | ✓ 1742ms | ✓ 1187ms | http |
| 45.167.124.52:8080 | ✓ 1714ms | ✓ 1702ms | ✓ 1169ms | 否 | ✓ 1339ms | http |
| 45.140.147.155:1081 | ✓ 864ms | 否 | ✓ 840ms | ✓ 1586ms | ✓ 1211ms | http |
| 8.219.97.248:80 | ✓ 1974ms | 否 | ✓ 1070ms | ✓ 1529ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1231ms | ✓ 1654ms | ✓ 1169ms | ✓ 1537ms | ✓ 1179ms | http |
| 138.124.53.25:7443 | ✓ 808ms | 否 | ✓ 1926ms | ✓ 1555ms | ✓ 1497ms | http |
| 81.70.169.194:80 | ✓ 1229ms | ✓ 1530ms | ✓ 1161ms | ✓ 1484ms | ✓ 1195ms | http |
| 101.43.255.96:80 | ✓ 1131ms | ✓ 1489ms | ✓ 1178ms | ✓ 1543ms | ✓ 1223ms | http |
| 150.230.249.50:1080 | ✓ 1496ms | ✓ 1812ms | 否 | ✓ 1262ms | 否 | http |
| 14.225.212.37:7890 | ✓ 1077ms | 否 | ✓ 1241ms | ✓ 1414ms | ✓ 1936ms | http |
| 38.145.208.193:8443 | ✓ 757ms | ✓ 1029ms | ✓ 347ms | ✓ 1187ms | ✓ 681ms | http |
| 38.145.208.192:8443 | ✓ 758ms | ✓ 1006ms | ✓ 368ms | ✓ 1171ms | ✓ 714ms | http |
| 38.145.208.189:8443 | ✓ 757ms | ✓ 1030ms | ✓ 346ms | ✓ 1186ms | ✓ 682ms | http |
| 38.145.208.191:8443 | ✓ 756ms | ✓ 1024ms | ✓ 351ms | ✓ 1197ms | ✓ 682ms | http |
| 38.145.218.163:8443 | ✓ 780ms | ✓ 865ms | ✓ 486ms | ✓ 1035ms | ✓ 884ms | http |
| 45.136.130.250:8443 | ✓ 759ms | ✓ 882ms | ✓ 718ms | ✓ 1045ms | ✓ 852ms | http |
| 38.145.208.196:8443 | ✓ 757ms | ✓ 1567ms | ✓ 389ms | ✓ 846ms | ✓ 678ms | http |
| 45.136.130.248:8443 | ✓ 758ms | ✓ 879ms | ✓ 718ms | ✓ 1065ms | ✓ 831ms | http |
| 38.145.218.51:8443 | ✓ 758ms | ✓ 894ms | ✓ 695ms | ✓ 1054ms | ✓ 843ms | http |
| 45.136.130.246:8443 | ✓ 758ms | ✓ 843ms | ✓ 762ms | ✓ 1051ms | ✓ 849ms | http |
| 38.145.208.190:8443 | ✓ 757ms | ✓ 1565ms | ✓ 375ms | ✓ 882ms | ✓ 711ms | http |
| 38.145.203.43:8443 | ✓ 778ms | ✓ 920ms | ✓ 719ms | ✓ 1164ms | ✓ 679ms | http |
| 38.145.203.46:8443 | ✓ 775ms | ✓ 922ms | ✓ 725ms | ✓ 1170ms | ✓ 674ms | http |
| 38.145.218.13:8443 | ✓ 758ms | ✓ 876ms | ✓ 724ms | ✓ 1077ms | ✓ 832ms | http |
| 38.145.203.35:8443 | ✓ 777ms | ✓ 879ms | ✓ 773ms | ✓ 1151ms | ✓ 680ms | http |
| 38.145.218.113:8443 | ✓ 777ms | ✓ 1485ms | ✓ 251ms | ✓ 1067ms | ✓ 698ms | http |
| 38.145.203.32:8443 | ✓ 778ms | ✓ 1483ms | ✓ 284ms | ✓ 1015ms | ✓ 734ms | http |
| 38.145.203.132:8443 | ✓ 776ms | ✓ 838ms | ✓ 963ms | ✓ 1062ms | ✓ 684ms | http |
| 38.145.203.19:8443 | ✓ 777ms | ✓ 926ms | ✓ 722ms | ✓ 1171ms | ✓ 725ms | http |
| 38.145.208.235:8443 | ✓ 778ms | ✓ 843ms | ✓ 942ms | ✓ 1061ms | ✓ 711ms | http |
| 38.145.208.239:8443 | ✓ 776ms | ✓ 848ms | ✓ 954ms | ✓ 1042ms | ✓ 726ms | http |
| 38.145.220.168:8443 | ✓ 777ms | ✓ 1149ms | ✓ 745ms | ✓ 1025ms | ✓ 680ms | http |
| 38.145.208.228:8443 | ✓ 776ms | ✓ 841ms | ✓ 977ms | ✓ 1054ms | ✓ 697ms | http |
| 38.145.220.103:8443 | ✓ 778ms | ✓ 1136ms | ✓ 827ms | ✓ 939ms | ✓ 710ms | http |
| 38.34.183.16:8443 | ✓ 758ms | ✓ 970ms | ✓ 922ms | ✓ 895ms | ✓ 837ms | http |
| 38.145.203.96:8443 | ✓ 779ms | ✓ 1141ms | ✓ 750ms | ✓ 1014ms | ✓ 696ms | http |
| 38.145.220.22:8443 | ✓ 768ms | ✓ 996ms | ✓ 877ms | ✓ 896ms | ✓ 857ms | http |
| 38.145.220.8:8443 | ✓ 758ms | ✓ 950ms | ✓ 944ms | ✓ 883ms | ✓ 847ms | http |
| 38.145.220.49:8443 | ✓ 758ms | ✓ 991ms | ✓ 892ms | ✓ 884ms | ✓ 866ms | http |
| 38.145.218.210:8443 | ✓ 778ms | ✓ 1666ms | ✓ 298ms | ✓ 958ms | ✓ 690ms | http |
| 38.145.220.20:8443 | ✓ 758ms | ✓ 1591ms | ✓ 377ms | ✓ 929ms | ✓ 738ms | http |
| 38.145.220.102:8443 | ✓ 777ms | ✓ 1118ms | ✓ 769ms | ✓ 1072ms | ✓ 689ms | http |
| 45.136.131.35:8443 | ✓ 759ms | ✓ 1023ms | ✓ 923ms | ✓ 881ms | ✓ 681ms | http |
| 38.145.220.6:8443 | ✓ 758ms | ✓ 974ms | ✓ 990ms | ✓ 936ms | ✓ 793ms | http |
| 45.136.131.32:8443 | ✓ 762ms | ✓ 1022ms | ✓ 936ms | ✓ 865ms | ✓ 714ms | http |
| 38.145.220.9:8443 | ✓ 760ms | ✓ 988ms | ✓ 899ms | ✓ 994ms | ✓ 833ms | http |
| 38.145.220.82:8443 | ✓ 758ms | ✓ 1039ms | ✓ 764ms | ✓ 1141ms | ✓ 813ms | http |
| 38.145.220.79:8443 | ✓ 758ms | ✓ 1033ms | ✓ 768ms | ✓ 1136ms | ✓ 816ms | http |
| 38.145.220.39:8443 | ✓ 759ms | ✓ 1054ms | ✓ 726ms | ✓ 1139ms | ✓ 844ms | http |
| 38.145.218.206:8443 | ✓ 759ms | ✓ 1016ms | ✓ 782ms | ✓ 1122ms | ✓ 844ms | http |
| 38.145.220.65:8443 | ✓ 757ms | ✓ 1058ms | ✓ 723ms | ✓ 1155ms | ✓ 844ms | http |
| 38.145.208.201:8447 | ✓ 782ms | ✓ 886ms | ✓ 808ms | ✓ 1173ms | ✓ 866ms | http |
| 38.145.220.81:8443 | 否 | ✓ 845ms | ✓ 271ms | ✓ 908ms | ✓ 694ms | http |
| 45.136.130.211:8447 | ✓ 779ms | 否 | ✓ 282ms | ✓ 886ms | ✓ 787ms | http |
| 38.145.220.11:8443 | 否 | ✓ 1052ms | ✓ 283ms | ✓ 871ms | ✓ 713ms | http |
| 38.145.208.209:8443 | ✓ 759ms | ✓ 1086ms | ✓ 957ms | ✓ 983ms | ✓ 723ms | http |
| 38.145.208.203:8443 | ✓ 760ms | ✓ 1079ms | ✓ 913ms | ✓ 1048ms | ✓ 712ms | http |
| 38.145.208.208:8443 | ✓ 759ms | ✓ 1079ms | ✓ 971ms | ✓ 979ms | ✓ 700ms | http |
| 38.145.208.210:8443 | ✓ 758ms | ✓ 1073ms | ✓ 924ms | ✓ 1008ms | ✓ 738ms | http |
| 38.145.208.206:8443 | ✓ 759ms | ✓ 1085ms | ✓ 920ms | ✓ 1036ms | ✓ 724ms | http |
| 38.145.208.214:8443 | ✓ 758ms | ✓ 1127ms | ✓ 886ms | ✓ 1042ms | ✓ 738ms | http |
| 38.145.208.207:8443 | ✓ 758ms | ✓ 1125ms | ✓ 868ms | ✓ 1052ms | ✓ 725ms | http |
| 38.145.208.205:8443 | ✓ 759ms | ✓ 1087ms | ✓ 912ms | ✓ 1027ms | ✓ 750ms | http |
| 192.71.213.85:9091 | ✓ 1258ms | 否 | ✓ 1516ms | ✓ 1926ms | 否 | http |
| 171.251.173.39:5104 | ✓ 1572ms | 否 | ✓ 1580ms | ✓ 1938ms | ✓ 1704ms | http |
| 165.227.5.10:8888 | ✓ 567ms | 否 | ✓ 715ms | ✓ 1455ms | ✓ 752ms | http |
| 185.115.74.185:8080 | ✓ 691ms | ✓ 1930ms | ✓ 1249ms | 否 | 否 | http |
| 45.140.147.82:1082 | ✓ 548ms | ✓ 1914ms | ✓ 1448ms | 否 | ✓ 919ms | http |
| 13.36.243.194:9899 | ✓ 1225ms | 否 | ✓ 1585ms | 否 | ✓ 1711ms | http |
| 45.140.147.82:1081 | ✓ 547ms | 否 | ✓ 1362ms | ✓ 1696ms | ✓ 1110ms | http |
| 47.105.98.23:3128 | ✓ 1278ms | ✓ 1593ms | ✓ 1038ms | ✓ 1514ms | ✓ 1153ms | http |
| 116.80.95.238:7777 | ✓ 1643ms | 否 | ✓ 1820ms | 否 | ✓ 1855ms | http |
| 120.92.212.16:7890 | ✓ 1499ms | 否 | ✓ 1225ms | ✓ 1426ms | 否 | http |
| 62.60.177.204:34094 | ✓ 465ms | ✓ 1846ms | ✓ 165ms | ✓ 878ms | ✓ 738ms | http |
| 45.136.130.223:8443 | 否 | ✓ 840ms | ✓ 706ms | ✓ 900ms | ✓ 674ms | http |
| 129.213.139.179:8080 | ✓ 496ms | ✓ 1538ms | ✓ 1161ms | 否 | ✓ 1323ms | http |
| 38.145.220.245:8443 | ✓ 731ms | ✓ 1007ms | ✓ 248ms | ✓ 1055ms | ✓ 820ms | http |
| 38.145.218.101:8447 | ✓ 732ms | ✓ 959ms | ✓ 290ms | ✓ 1035ms | ✓ 823ms | http |
| 45.136.130.220:8443 | ✓ 733ms | ✓ 984ms | ✓ 260ms | ✓ 1077ms | ✓ 810ms | http |
| 45.136.131.30:8447 | ✓ 734ms | ✓ 926ms | ✓ 741ms | ✓ 872ms | ✓ 719ms | http |
| 45.136.131.28:8447 | ✓ 734ms | ✓ 873ms | ✓ 784ms | ✓ 873ms | ✓ 762ms | http |
| 38.145.203.135:8443 | ✓ 734ms | ✓ 1400ms | ✓ 284ms | ✓ 947ms | ✓ 718ms | http |
| 38.145.218.134:8443 | ✓ 732ms | ✓ 1011ms | ✓ 292ms | ✓ 964ms | ✓ 1006ms | http |
| 38.145.218.162:8443 | ✓ 733ms | ✓ 997ms | ✓ 274ms | ✓ 996ms | ✓ 983ms | http |
| 38.145.218.161:8443 | ✓ 733ms | ✓ 979ms | ✓ 305ms | ✓ 1019ms | ✓ 1006ms | http |
| 45.136.131.42:8447 | ✓ 732ms | ✓ 981ms | ✓ 838ms | ✓ 953ms | ✓ 695ms | http |
| 45.136.130.253:8443 | ✓ 733ms | ✓ 945ms | ✓ 278ms | ✓ 1067ms | ✓ 699ms | http |
| 45.136.130.249:8443 | ✓ 734ms | ✓ 965ms | ✓ 255ms | ✓ 1050ms | ✓ 699ms | http |
| 45.136.130.252:8443 | ✓ 733ms | ✓ 994ms | ✓ 246ms | ✓ 1028ms | ✓ 695ms | http |
| 45.136.130.247:8443 | ✓ 733ms | ✓ 949ms | ✓ 262ms | ✓ 1061ms | ✓ 680ms | http |
| 38.145.218.76:8443 | ✓ 733ms | ✓ 944ms | ✓ 274ms | ✓ 1048ms | ✓ 688ms | http |
| 216.180.127.45:1080 | ✓ 453ms | ✓ 1801ms | ✓ 1068ms | 否 | ✓ 854ms | http |
| 59.46.216.131:30001 | ✓ 1137ms | ✓ 1526ms | ✓ 1279ms | 否 | ✓ 1332ms | http |
| 38.145.203.45:8443 | ✓ 755ms | ✓ 1051ms | ✓ 784ms | ✓ 1571ms | 否 | http |
| 38.145.203.41:8443 | ✓ 756ms | ✓ 1089ms | ✓ 737ms | ✓ 1600ms | 否 | http |
| 38.145.203.39:8443 | ✓ 755ms | ✓ 1069ms | ✓ 780ms | ✓ 1594ms | 否 | http |
| 38.145.203.34:8443 | ✓ 755ms | ✓ 1044ms | ✓ 791ms | ✓ 1629ms | 否 | http |
| 178.236.245.59:3128 | ✓ 1594ms | 否 | ✓ 656ms | ✓ 1575ms | 否 | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
