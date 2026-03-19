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

最后更新：2026-03-19 05:50:12 UTC（2026-03-19 13:50:12 UTC+8）

**代理总数：165**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 165 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 165 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 202.155.12.161:443 | ✓ 1413ms | 否 | ✓ 687ms | ✓ 973ms | ✓ 947ms | http |
| 147.161.210.140:8800 | ✓ 1414ms | 否 | ✓ 738ms | ✓ 734ms | ✓ 706ms | http |
| 113.160.132.26:8080 | ✓ 1604ms | ✓ 1306ms | ✓ 945ms | ✓ 1230ms | ✓ 946ms | http |
| 101.47.73.135:3128 | ✓ 796ms | 否 | ✓ 1593ms | ✓ 1088ms | ✓ 1026ms | http |
| 147.161.239.240:8800 | ✓ 1396ms | 否 | ✓ 1880ms | ✓ 1887ms | ✓ 1670ms | http |
| 85.198.96.242:3128 | ✓ 1427ms | 否 | 否 | ✓ 1925ms | ✓ 1488ms | http |
| 47.77.193.180:1080 | ✓ 468ms | ✓ 1856ms | ✓ 436ms | ✓ 670ms | ✓ 500ms | http |
| 35.225.22.61:80 | ✓ 478ms | 否 | ✓ 1000ms | ✓ 1436ms | 否 | http |
| 120.92.212.16:8890 | ✓ 946ms | ✓ 1197ms | ✓ 1158ms | ✓ 1237ms | 否 | http |
| 45.125.67.37:443 | ✓ 1496ms | 否 | ✓ 1255ms | ✓ 1042ms | ✓ 1096ms | http |
| 1.231.81.166:3128 | ✓ 930ms | ✓ 1168ms | 否 | ✓ 1549ms | ✓ 1146ms | http |
| 115.231.181.40:8128 | ✓ 865ms | 否 | ✓ 972ms | ✓ 1167ms | ✓ 922ms | http |
| 212.192.12.90:6005 | 否 | 否 | ✓ 1109ms | ✓ 1177ms | ✓ 1002ms | http |
| 160.238.65.4:3128 | ✓ 1390ms | ✓ 1709ms | ✓ 1765ms | 否 | ✓ 1666ms | http |
| 160.238.65.8:3128 | ✓ 1391ms | ✓ 1852ms | ✓ 1524ms | 否 | ✓ 1765ms | http |
| 160.238.65.9:3128 | ✓ 1384ms | ✓ 1788ms | ✓ 1594ms | 否 | ✓ 1749ms | http |
| 160.238.65.6:3128 | ✓ 1391ms | 否 | ✓ 1375ms | 否 | ✓ 1756ms | http |
| 160.238.65.2:3128 | ✓ 1696ms | 否 | ✓ 1076ms | 否 | ✓ 1761ms | http |
| 160.238.65.3:3128 | ✓ 1392ms | 否 | ✓ 1375ms | 否 | ✓ 1764ms | http |
| 160.238.65.5:3128 | ✓ 1690ms | 否 | ✓ 1650ms | 否 | ✓ 1197ms | http |
| 160.238.65.7:3128 | ✓ 1384ms | ✓ 1896ms | ✓ 1488ms | 否 | ✓ 1752ms | http |
| 165.227.5.10:8888 | ✓ 1224ms | ✓ 1074ms | ✓ 1566ms | ✓ 1104ms | 否 | http |
| 38.145.220.175:8443 | ✓ 803ms | 否 | ✓ 218ms | ✓ 678ms | ✓ 720ms | http |
| 174.138.24.77:1080 | ✓ 712ms | 否 | ✓ 924ms | ✓ 989ms | ✓ 773ms | http |
| 38.145.220.81:8452 | ✓ 322ms | ✓ 708ms | ✓ 1539ms | ✓ 1772ms | ✓ 579ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 953ms | ✓ 1218ms | ✓ 935ms | http |
| 101.43.127.100:8877 | ✓ 862ms | ✓ 1855ms | ✓ 1315ms | ✓ 1307ms | ✓ 825ms | http |
| 121.237.181.137:8888 | ✓ 930ms | ✓ 1379ms | ✓ 1571ms | 否 | ✓ 939ms | http |
| 160.250.4.13:1 | ✓ 1818ms | 否 | ✓ 1732ms | ✓ 1453ms | ✓ 1216ms | http |
| 38.34.179.176:8443 | ✓ 660ms | ✓ 619ms | ✓ 106ms | ✓ 819ms | ✓ 526ms | http |
| 38.145.208.240:8448 | ✓ 626ms | ✓ 657ms | ✓ 1116ms | ✓ 793ms | ✓ 524ms | http |
| 34.101.184.164:3128 | ✓ 1692ms | 否 | ✓ 1304ms | 否 | ✓ 1517ms | http |
| 38.145.220.182:8445 | 否 | ✓ 865ms | ✓ 787ms | ✓ 754ms | ✓ 830ms | http |
| 45.136.131.60:8452 | 否 | ✓ 620ms | ✓ 1619ms | ✓ 960ms | ✓ 514ms | http |
| 38.145.218.161:8444 | 否 | 否 | ✓ 482ms | ✓ 940ms | ✓ 592ms | http |
| 38.34.179.156:8450 | 否 | ✓ 1047ms | ✓ 1430ms | ✓ 1452ms | ✓ 645ms | http |
| 38.145.220.27:8451 | 否 | ✓ 1748ms | ✓ 634ms | ✓ 1497ms | ✓ 516ms | http |
| 38.34.179.47:8448 | 否 | ✓ 1960ms | ✓ 1003ms | ✓ 1615ms | ✓ 510ms | http |
| 168.235.110.63:3128 | ✓ 1027ms | 否 | ✓ 1253ms | ✓ 1422ms | ✓ 991ms | http |
| 137.220.150.170:6005 | ✓ 988ms | 否 | ✓ 786ms | ✓ 1103ms | ✓ 851ms | http |
| 185.191.236.162:3128 | ✓ 773ms | 否 | ✓ 1052ms | ✓ 1766ms | ✓ 1196ms | http |
| 185.104.249.25:3128 | ✓ 809ms | 否 | ✓ 810ms | 否 | ✓ 1737ms | http |
| 38.145.208.183:8450 | 否 | ✓ 1382ms | ✓ 1120ms | 否 | ✓ 797ms | http |
| 148.153.56.51:80 | ✓ 600ms | ✓ 599ms | ✓ 785ms | ✓ 857ms | ✓ 814ms | http |
| 38.34.179.175:8443 | ✓ 613ms | ✓ 655ms | ✓ 253ms | ✓ 690ms | ✓ 1771ms | http |
| 35.194.4.51:3128 | ✓ 635ms | ✓ 1380ms | ✓ 849ms | ✓ 1150ms | ✓ 1079ms | http |
| 146.190.232.76:3128 | ✓ 697ms | 否 | 否 | ✓ 1616ms | ✓ 1494ms | http |
| 137.220.150.152:6005 | ✓ 719ms | 否 | ✓ 785ms | ✓ 1208ms | ✓ 1099ms | http |
| 23.224.193.42:3128 | 否 | ✓ 1738ms | ✓ 1006ms | ✓ 1124ms | ✓ 1629ms | http |
| 38.34.179.60:8450 | ✓ 541ms | ✓ 1598ms | ✓ 865ms | ✓ 831ms | ✓ 538ms | http |
| 38.34.179.89:8443 | ✓ 131ms | ✓ 1276ms | ✓ 810ms | ✓ 702ms | ✓ 542ms | http |
| 38.34.179.88:8443 | ✓ 797ms | ✓ 1013ms | ✓ 106ms | 否 | 否 | http |
| 38.145.220.182:8443 | ✓ 649ms | ✓ 686ms | ✓ 76ms | ✓ 737ms | ✓ 575ms | http |
| 38.34.179.91:8443 | ✓ 125ms | ✓ 871ms | ✓ 105ms | ✓ 663ms | ✓ 526ms | http |
| 45.136.131.54:8448 | ✓ 124ms | ✓ 759ms | ✓ 179ms | ✓ 668ms | ✓ 515ms | http |
| 20.78.213.56:80 | ✓ 986ms | 否 | 否 | ✓ 997ms | ✓ 1302ms | http |
| 133.242.138.34:8100 | ✓ 1010ms | 否 | ✓ 1682ms | ✓ 1911ms | ✓ 1532ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 1080ms | ✓ 1509ms | ✓ 1792ms | http |
| 194.5.212.40:8080 | ✓ 1178ms | 否 | ✓ 1522ms | 否 | ✓ 1625ms | http |
| 45.136.131.47:8447 | ✓ 901ms | ✓ 863ms | ✓ 84ms | ✓ 690ms | ✓ 1838ms | http |
| 45.136.131.53:8447 | ✓ 888ms | ✓ 1903ms | 否 | ✓ 1215ms | ✓ 576ms | http |
| 45.136.131.49:8447 | ✓ 891ms | ✓ 1896ms | 否 | ✓ 1231ms | ✓ 597ms | http |
| 45.136.131.48:8447 | ✓ 899ms | ✓ 1992ms | 否 | ✓ 1138ms | ✓ 588ms | http |
| 137.220.151.110:6005 | ✓ 1312ms | 否 | ✓ 880ms | ✓ 1515ms | ✓ 1009ms | http |
| 209.126.10.139:3128 | ✓ 1487ms | 否 | ✓ 986ms | ✓ 1114ms | ✓ 1050ms | http |
| 121.230.8.91:1080 | 否 | 否 | ✓ 1880ms | ✓ 1453ms | ✓ 1127ms | http |
| 45.136.131.48:8451 | 否 | ✓ 1094ms | ✓ 734ms | ✓ 669ms | ✓ 678ms | http |
| 38.145.220.79:8448 | ✓ 1390ms | ✓ 1447ms | ✓ 851ms | ✓ 966ms | ✓ 653ms | http |
| 38.145.220.100:8443 | ✓ 667ms | ✓ 610ms | ✓ 1108ms | ✓ 863ms | ✓ 544ms | http |
| 137.220.150.104:6005 | ✓ 1715ms | 否 | ✓ 1743ms | ✓ 1122ms | ✓ 1390ms | http |
| 38.145.220.102:8443 | ✓ 229ms | ✓ 609ms | ✓ 82ms | ✓ 658ms | ✓ 500ms | http |
| 38.145.220.56:8443 | ✓ 164ms | ✓ 644ms | ✓ 83ms | ✓ 653ms | ✓ 505ms | http |
| 38.145.220.103:8443 | ✓ 227ms | ✓ 646ms | ✓ 74ms | ✓ 666ms | ✓ 490ms | http |
| 38.145.220.196:8443 | ✓ 189ms | ✓ 682ms | ✓ 73ms | ✓ 665ms | ✓ 504ms | http |
| 38.145.220.198:8448 | ✓ 184ms | ✓ 1016ms | ✓ 78ms | ✓ 651ms | ✓ 502ms | http |
| 38.145.203.87:8443 | ✓ 298ms | ✓ 1572ms | ✓ 72ms | ✓ 682ms | ✓ 501ms | http |
| 38.145.218.210:8443 | ✓ 228ms | ✓ 1091ms | ✓ 76ms | ✓ 844ms | ✓ 506ms | http |
| 38.145.220.72:8443 | ✓ 180ms | ✓ 1190ms | ✓ 89ms | ✓ 677ms | ✓ 517ms | http |
| 38.145.220.93:8443 | ✓ 167ms | ✓ 1180ms | ✓ 87ms | ✓ 683ms | ✓ 502ms | http |
| 38.145.208.215:8451 | ✓ 968ms | ✓ 1657ms | ✓ 99ms | ✓ 657ms | ✓ 518ms | http |
| 38.145.203.86:8443 | ✓ 124ms | ✓ 1989ms | ✓ 213ms | ✓ 651ms | ✓ 502ms | http |
| 38.145.220.193:8443 | ✓ 194ms | ✓ 1459ms | ✓ 88ms | ✓ 665ms | ✓ 505ms | http |
| 38.145.203.96:8443 | ✓ 282ms | ✓ 1032ms | ✓ 85ms | ✓ 666ms | ✓ 1546ms | http |
| 38.145.220.96:8443 | ✓ 167ms | ✓ 1648ms | ✓ 80ms | ✓ 651ms | ✓ 540ms | http |
| 38.34.179.83:8448 | ✓ 127ms | ✓ 1401ms | ✓ 82ms | ✓ 694ms | ✓ 519ms | http |
| 45.136.131.52:8443 | ✓ 241ms | ✓ 1359ms | ✓ 74ms | ✓ 650ms | ✓ 491ms | http |
| 45.136.131.56:8447 | ✓ 231ms | ✓ 1360ms | ✓ 77ms | ✓ 647ms | ✓ 503ms | http |
| 38.145.220.179:8443 | ✓ 203ms | ✓ 1444ms | ✓ 84ms | ✓ 681ms | ✓ 1163ms | http |
| 103.113.70.189:1081 | ✓ 361ms | ✓ 1176ms | ✓ 892ms | ✓ 1182ms | ✓ 961ms | http |
| 38.34.179.178:8445 | ✓ 201ms | ✓ 664ms | ✓ 634ms | ✓ 1279ms | ✓ 737ms | http |
| 45.136.131.50:8443 | ✓ 244ms | ✓ 1746ms | ✓ 75ms | ✓ 665ms | ✓ 500ms | http |
| 45.136.131.57:8443 | ✓ 234ms | ✓ 1736ms | ✓ 87ms | ✓ 665ms | ✓ 505ms | http |
| 38.34.179.71:8447 | ✓ 89ms | ✓ 708ms | ✓ 187ms | ✓ 879ms | 否 | http |
| 38.145.208.181:8445 | ✓ 337ms | ✓ 622ms | ✓ 737ms | ✓ 1032ms | 否 | http |
| 38.34.179.85:8443 | ✓ 123ms | ✓ 1934ms | ✓ 82ms | ✓ 674ms | ✓ 1507ms | http |
| 38.145.218.163:8451 | ✓ 1551ms | ✓ 619ms | ✓ 115ms | ✓ 735ms | ✓ 709ms | http |
| 38.34.179.14:8450 | ✓ 1756ms | ✓ 1148ms | ✓ 118ms | ✓ 929ms | 否 | http |
| 38.34.183.233:8448 | ✓ 478ms | ✓ 1397ms | ✓ 657ms | ✓ 928ms | ✓ 884ms | http |
| 38.34.183.225:8450 | ✓ 1293ms | ✓ 1670ms | ✓ 781ms | 否 | ✓ 1127ms | http |
| 38.145.220.33:8448 | 否 | 否 | ✓ 443ms | ✓ 937ms | ✓ 1210ms | http |

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
