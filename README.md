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

最后更新：2026-03-15 21:29:54 UTC（2026-03-16 05:29:54 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 325ms | ✓ 967ms | ✓ 1053ms | ✓ 1156ms | ✓ 781ms | http |
| 137.220.151.110:6005 | ✓ 1074ms | 否 | ✓ 883ms | ✓ 1309ms | ✓ 988ms | http |
| 5.129.206.247:8888 | ✓ 1077ms | ✓ 1886ms | ✓ 1371ms | 否 | ✓ 1842ms | http |
| 194.87.43.49:8888 | ✓ 1016ms | 否 | ✓ 1566ms | 否 | ✓ 1793ms | http |
| 113.160.132.26:8080 | ✓ 1561ms | ✓ 1479ms | ✓ 1375ms | ✓ 1323ms | ✓ 1639ms | http |
| 137.220.150.152:6005 | ✓ 903ms | 否 | ✓ 1055ms | ✓ 1351ms | ✓ 983ms | http |
| 45.167.124.52:8080 | ✓ 1379ms | 否 | ✓ 1259ms | ✓ 1636ms | ✓ 1305ms | http |
| 62.60.177.204:34094 | ✓ 467ms | ✓ 1255ms | ✓ 1221ms | ✓ 1123ms | ✓ 879ms | http |
| 103.84.95.54:7890 | ✓ 816ms | 否 | ✓ 1139ms | ✓ 1211ms | ✓ 766ms | http |
| 152.70.16.146:8888 | ✓ 968ms | 否 | ✓ 1895ms | ✓ 1975ms | ✓ 1407ms | http |
| 59.46.216.131:30001 | ✓ 1091ms | ✓ 1550ms | ✓ 1258ms | ✓ 1478ms | ✓ 1146ms | http |
| 47.79.40.38:55000 | ✓ 1095ms | ✓ 1558ms | ✓ 1096ms | ✓ 1089ms | ✓ 758ms | http |
| 38.145.218.82:8443 | ✓ 754ms | ✓ 948ms | ✓ 1487ms | ✓ 999ms | ✓ 734ms | http |
| 160.238.65.4:3128 | ✓ 688ms | 否 | ✓ 1431ms | ✓ 1276ms | 否 | http |
| 160.238.65.7:3128 | ✓ 694ms | 否 | ✓ 1426ms | ✓ 1288ms | 否 | http |
| 160.238.65.9:3128 | ✓ 661ms | 否 | ✓ 1459ms | ✓ 1290ms | 否 | http |
| 160.238.65.2:3128 | ✓ 681ms | 否 | ✓ 1439ms | ✓ 1302ms | 否 | http |
| 160.238.65.8:3128 | ✓ 687ms | ✓ 1862ms | ✓ 1571ms | ✓ 1303ms | 否 | http |
| 160.238.65.5:3128 | ✓ 658ms | ✓ 1787ms | ✓ 1673ms | ✓ 1314ms | 否 | http |
| 160.238.65.6:3128 | ✓ 899ms | 否 | ✓ 1221ms | ✓ 1317ms | 否 | http |
| 149.50.116.240:1080 | ✓ 578ms | 否 | ✓ 1249ms | ✓ 1721ms | 否 | http |
| 160.238.65.3:3128 | ✓ 680ms | ✓ 1766ms | ✓ 1673ms | ✓ 1276ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1045ms | ✓ 1268ms | ✓ 1041ms | ✓ 1321ms | ✓ 1034ms | http |
| 38.145.203.135:8443 | ✓ 771ms | ✓ 1186ms | ✓ 1228ms | ✓ 1156ms | ✓ 1046ms | http |
| 172.212.68.37:3128 | ✓ 284ms | ✓ 1435ms | ✓ 1057ms | ✓ 1171ms | ✓ 766ms | http |
| 101.43.127.100:8877 | ✓ 968ms | ✓ 1163ms | ✓ 1038ms | ✓ 1209ms | ✓ 998ms | http |
| 120.92.212.16:7890 | ✓ 1062ms | ✓ 1340ms | ✓ 1210ms | ✓ 1324ms | ✓ 1099ms | http |
| 81.70.169.194:80 | ✓ 1172ms | ✓ 1420ms | ✓ 1135ms | ✓ 1366ms | ✓ 1210ms | http |
| 101.43.255.96:80 | ✓ 1142ms | ✓ 1462ms | ✓ 1094ms | ✓ 1459ms | ✓ 1152ms | http |
| 116.80.49.156:3172 | ✓ 1605ms | 否 | ✓ 1637ms | ✓ 1963ms | 否 | http |
| 86.53.183.16:1080 | ✓ 840ms | ✓ 1713ms | ✓ 1351ms | ✓ 1837ms | ✓ 1626ms | http |
| 104.129.202.127:12354 | 否 | 否 | ✓ 1448ms | ✓ 1387ms | ✓ 926ms | http |
| 185.115.74.185:8080 | ✓ 820ms | ✓ 1683ms | ✓ 1565ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 592ms | 否 | ✓ 728ms | 否 | ✓ 813ms | http |
| 72.11.150.178:6005 | ✓ 555ms | ✓ 1970ms | ✓ 1011ms | ✓ 1193ms | ✓ 934ms | http |
| 137.220.150.104:6005 | ✓ 879ms | 否 | ✓ 895ms | ✓ 1302ms | ✓ 1020ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1526ms | ✓ 1404ms | ✓ 1968ms | http |
| 85.198.96.242:3128 | ✓ 1787ms | 否 | ✓ 1615ms | ✓ 1711ms | ✓ 1309ms | http |
| 212.192.13.76:6005 | 否 | 否 | ✓ 1631ms | ✓ 1330ms | ✓ 1884ms | http |
| 137.220.128.173:8866 | ✓ 1691ms | 否 | ✓ 1448ms | ✓ 1785ms | ✓ 1518ms | http |
| 137.220.128.149:8866 | ✓ 1362ms | 否 | ✓ 1783ms | ✓ 1788ms | ✓ 1547ms | http |
| 185.221.215.134:7788 | ✓ 1533ms | 否 | ✓ 1658ms | 否 | ✓ 1689ms | http |
| 14.225.212.37:7890 | 否 | 否 | ✓ 1374ms | ✓ 1614ms | ✓ 1510ms | http |
| 95.3.9.78:3128 | ✓ 806ms | ✓ 1705ms | ✓ 738ms | ✓ 1680ms | ✓ 1289ms | http |
| 95.3.9.78:8080 | ✓ 912ms | 否 | ✓ 1080ms | ✓ 1615ms | 否 | http |
| 168.235.110.63:3128 | ✓ 225ms | ✓ 1974ms | ✓ 1819ms | ✓ 1110ms | ✓ 748ms | http |
| 178.236.245.17:3128 | ✓ 591ms | 否 | ✓ 1174ms | 否 | ✓ 1640ms | http |
| 178.236.245.59:3128 | ✓ 776ms | 否 | ✓ 1004ms | ✓ 1882ms | 否 | http |
| 104.129.202.127:10810 | ✓ 1347ms | 否 | ✓ 994ms | ✓ 1016ms | ✓ 824ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1771ms | ✓ 1349ms | ✓ 1573ms | 否 | http |
| 165.227.5.10:8888 | ✓ 860ms | ✓ 1647ms | ✓ 224ms | ✓ 982ms | ✓ 1224ms | http |
| 103.113.70.189:1081 | ✓ 346ms | ✓ 892ms | 否 | ✓ 1001ms | ✓ 821ms | http |
| 178.156.224.42:3128 | ✓ 890ms | ✓ 1902ms | 否 | 否 | ✓ 1920ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1656ms | ✓ 1358ms | ✓ 1660ms | http |
| 8.219.97.248:80 | ✓ 1932ms | 否 | 否 | ✓ 1406ms | ✓ 1452ms | http |
| 45.140.147.155:1082 | ✓ 1019ms | 否 | ✓ 1506ms | ✓ 1608ms | ✓ 1864ms | http |
| 45.140.147.155:1081 | ✓ 410ms | 否 | ✓ 888ms | ✓ 1628ms | ✓ 1167ms | http |
| 106.117.208.101:7890 | ✓ 1135ms | ✓ 1511ms | ✓ 1303ms | 否 | ✓ 1228ms | http |
| 211.171.114.154:3128 | ✓ 1405ms | ✓ 1157ms | ✓ 1783ms | ✓ 1310ms | ✓ 1001ms | http |
| 193.23.200.251:10808 | ✓ 1075ms | ✓ 1991ms | ✓ 884ms | ✓ 1947ms | 否 | http |
| 45.168.238.193:8443 | ✓ 697ms | ✓ 1238ms | ✓ 680ms | ✓ 1413ms | ✓ 912ms | http |
| 198.24.188.138:35236 | ✓ 1186ms | 否 | ✓ 608ms | ✓ 1795ms | 否 | http |
| 129.213.162.27:17777 | ✓ 1038ms | ✓ 1624ms | ✓ 1818ms | ✓ 1836ms | ✓ 1141ms | http |
| 186.148.180.46:999 | ✓ 1062ms | ✓ 1763ms | ✓ 1226ms | ✓ 1852ms | ✓ 1340ms | http |
| 45.129.141.143:3128 | ✓ 1336ms | 否 | ✓ 1717ms | 否 | ✓ 1608ms | http |
| 181.78.44.63:999 | ✓ 992ms | ✓ 1899ms | ✓ 1358ms | ✓ 1974ms | ✓ 1873ms | http |
| 103.39.51.190:8080 | ✓ 1841ms | 否 | ✓ 1936ms | 否 | ✓ 1832ms | http |
| 45.136.198.40:3128 | ✓ 963ms | ✓ 1825ms | ✓ 1725ms | ✓ 1957ms | ✓ 1732ms | http |
| 223.16.170.103:80 | 否 | ✓ 1781ms | ✓ 1391ms | 否 | ✓ 1243ms | http |
| 45.119.85.216:3128 | 否 | 否 | ✓ 1229ms | ✓ 1829ms | ✓ 1408ms | http |
| 192.71.213.85:9091 | ✓ 1006ms | 否 | ✓ 761ms | ✓ 1895ms | 否 | http |
| 192.71.213.85:5678 | ✓ 1442ms | 否 | ✓ 1874ms | ✓ 1897ms | 否 | http |
| 121.230.9.160:1080 | ✓ 1707ms | ✓ 1831ms | ✓ 1552ms | ✓ 1551ms | ✓ 1189ms | http |
| 192.71.213.85:9812 | ✓ 840ms | 否 | ✓ 485ms | ✓ 1383ms | 否 | http |
| 137.184.1.87:3128 | 否 | 否 | ✓ 1141ms | ✓ 944ms | ✓ 691ms | http |
| 103.139.138.194:3128 | ✓ 1857ms | 否 | ✓ 1266ms | ✓ 1505ms | ✓ 1288ms | http |
| 61.52.131.172:8443 | ✓ 947ms | ✓ 1266ms | ✓ 1044ms | ✓ 1329ms | ✓ 1042ms | http |
| 194.5.212.40:8080 | ✓ 630ms | ✓ 1979ms | ✓ 1516ms | 否 | ✓ 1497ms | http |
| 8.222.175.80:6128 | ✓ 1267ms | 否 | ✓ 1005ms | ✓ 1197ms | ✓ 930ms | http |
| 164.90.151.28:3128 | ✓ 464ms | ✓ 1759ms | ✓ 1020ms | ✓ 863ms | ✓ 672ms | http |
| 92.119.127.213:6005 | ✓ 1288ms | 否 | ✓ 1468ms | ✓ 1652ms | ✓ 1987ms | http |
| 67.223.241.33:43080 | 否 | ✓ 1143ms | ✓ 1619ms | ✓ 1224ms | 否 | http |

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
