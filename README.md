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

最后更新：2026-03-06 13:58:49 UTC（2026-03-06 21:58:49 UTC+8）

**代理总数：74**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1503ms | ✓ 1178ms | ✓ 923ms | ✓ 874ms | ✓ 670ms | http |
| 152.42.195.165:8888 | ✓ 899ms | 否 | ✓ 688ms | ✓ 1006ms | ✓ 810ms | http |
| 205.209.118.30:3138 | ✓ 574ms | 否 | 否 | ✓ 1374ms | ✓ 1053ms | http |
| 125.128.12.144:3128 | 否 | 否 | ✓ 1881ms | ✓ 1893ms | ✓ 822ms | http |
| 217.76.245.80:999 | ✓ 1645ms | ✓ 1648ms | ✓ 1347ms | ✓ 1508ms | ✓ 1257ms | http |
| 167.172.69.123:8080 | ✓ 764ms | 否 | ✓ 853ms | ✓ 1182ms | ✓ 869ms | http |
| 61.72.221.234:3128 | ✓ 1524ms | 否 | ✓ 1788ms | ✓ 989ms | ✓ 774ms | http |
| 121.128.121.54:3128 | 否 | 否 | ✓ 1703ms | ✓ 1992ms | ✓ 744ms | http |
| 120.92.212.16:7890 | ✓ 1037ms | 否 | ✓ 1177ms | ✓ 1884ms | 否 | http |
| 14.56.107.244:3128 | ✓ 1696ms | ✓ 1995ms | ✓ 1044ms | 否 | 否 | http |
| 20.27.14.220:8561 | ✓ 453ms | ✓ 828ms | ✓ 493ms | 否 | 否 | http |
| 125.128.12.14:3128 | 否 | ✓ 1254ms | 否 | ✓ 1441ms | ✓ 1010ms | http |
| 61.72.221.194:3128 | ✓ 873ms | 否 | ✓ 1600ms | 否 | ✓ 1863ms | http |
| 61.72.110.54:3128 | ✓ 1630ms | 否 | ✓ 1660ms | 否 | ✓ 1997ms | http |
| 115.231.181.40:8128 | ✓ 1217ms | ✓ 1707ms | ✓ 1129ms | ✓ 1484ms | ✓ 863ms | http |
| 101.43.255.96:80 | 否 | ✓ 1332ms | ✓ 1053ms | ✓ 1415ms | ✓ 1204ms | http |
| 91.193.240.157:9877 | ✓ 1490ms | 否 | ✓ 1359ms | 否 | ✓ 1739ms | http |
| 185.115.74.185:8080 | ✓ 1441ms | ✓ 1989ms | ✓ 1579ms | 否 | 否 | http |
| 45.136.198.40:3128 | ✓ 833ms | 否 | ✓ 831ms | ✓ 1745ms | ✓ 1398ms | http |
| 103.86.131.62:80 | ✓ 886ms | 否 | 否 | ✓ 1201ms | ✓ 966ms | http |
| 39.104.201.40:7890 | ✓ 868ms | ✓ 1483ms | 否 | ✓ 1155ms | ✓ 905ms | http |
| 103.82.23.118:5242 | ✓ 1866ms | ✓ 1887ms | ✓ 1275ms | ✓ 1402ms | ✓ 1142ms | http |
| 47.95.231.180:8084 | ✓ 1793ms | ✓ 1184ms | ✓ 873ms | ✓ 1143ms | ✓ 891ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1781ms | ✓ 1615ms | ✓ 1607ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1613ms | ✓ 1609ms | ✓ 1887ms | http |
| 20.210.76.175:8561 | 否 | ✓ 1526ms | ✓ 996ms | ✓ 1159ms | ✓ 809ms | http |
| 20.27.11.248:8561 | ✓ 1079ms | ✓ 1086ms | ✓ 464ms | ✓ 738ms | ✓ 628ms | http |
| 20.210.39.153:8561 | ✓ 1860ms | 否 | 否 | ✓ 1748ms | ✓ 1262ms | http |
| 20.78.26.206:8561 | ✓ 1864ms | 否 | 否 | ✓ 1751ms | ✓ 1268ms | http |
| 20.78.118.91:8561 | ✓ 1863ms | 否 | 否 | ✓ 1794ms | ✓ 1251ms | http |
| 62.113.119.14:8080 | ✓ 846ms | 否 | ✓ 1233ms | ✓ 1723ms | ✓ 1311ms | http |
| 20.27.15.111:8561 | 否 | ✓ 1083ms | ✓ 435ms | ✓ 781ms | ✓ 587ms | http |
| 20.27.15.49:8561 | 否 | 否 | ✓ 1532ms | ✓ 1908ms | ✓ 947ms | http |
| 20.210.76.178:8561 | 否 | 否 | ✓ 1547ms | ✓ 1902ms | ✓ 941ms | http |
| 20.210.76.104:8561 | 否 | 否 | ✓ 1712ms | ✓ 1915ms | ✓ 894ms | http |
| 120.92.212.16:8890 | ✓ 983ms | 否 | ✓ 1716ms | ✓ 1187ms | ✓ 928ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 741ms | ✓ 1077ms | ✓ 1903ms | http |
| 167.172.69.123:80 | ✓ 688ms | 否 | ✓ 875ms | ✓ 1008ms | ✓ 832ms | http |
| 159.223.42.219:3128 | ✓ 1467ms | 否 | ✓ 783ms | ✓ 1448ms | ✓ 1046ms | http |
| 61.72.110.94:3128 | 否 | ✓ 1552ms | ✓ 1215ms | ✓ 1033ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1402ms | 否 | 否 | ✓ 1665ms | ✓ 1923ms | http |
| 47.77.193.180:1080 | ✓ 620ms | ✓ 1317ms | ✓ 348ms | ✓ 638ms | ✓ 477ms | http |
| 116.80.82.219:3172 | ✓ 1554ms | 否 | 否 | ✓ 1824ms | ✓ 1727ms | http |
| 101.32.244.83:8080 | ✓ 1371ms | ✓ 1340ms | ✓ 893ms | ✓ 1369ms | ✓ 1162ms | http |
| 121.43.196.210:8222 | ✓ 926ms | ✓ 1033ms | ✓ 832ms | ✓ 1046ms | ✓ 850ms | http |
| 121.43.196.213:8222 | ✓ 940ms | ✓ 994ms | ✓ 851ms | ✓ 1106ms | ✓ 874ms | http |
| 114.55.226.123:10086 | ✓ 1351ms | 否 | ✓ 1010ms | ✓ 1233ms | ✓ 990ms | http |
| 61.76.95.217:40088 | 否 | 否 | ✓ 1531ms | ✓ 1455ms | ✓ 1134ms | http |
| 35.225.22.61:80 | ✓ 489ms | 否 | ✓ 436ms | ✓ 1130ms | ✓ 938ms | http |
| 103.84.95.54:7890 | ✓ 658ms | 否 | 否 | ✓ 771ms | ✓ 964ms | http |
| 103.215.36.88:16474 | ✓ 1631ms | 否 | ✓ 1054ms | 否 | ✓ 991ms | http |
| 14.225.222.247:7890 | 否 | ✓ 1584ms | ✓ 1094ms | 否 | ✓ 1046ms | http |
| 116.80.82.228:3172 | ✓ 1927ms | 否 | ✓ 1514ms | ✓ 1797ms | 否 | http |
| 207.254.71.62:8088 | ✓ 1458ms | 否 | ✓ 1921ms | 否 | ✓ 1739ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1978ms | 否 | ✓ 1742ms | ✓ 1994ms | http |
| 1.234.153.14:80 | ✓ 663ms | 否 | ✓ 792ms | ✓ 825ms | ✓ 651ms | http |
| 20.78.213.56:80 | ✓ 788ms | ✓ 1730ms | ✓ 751ms | ✓ 1484ms | ✓ 761ms | http |
| 103.82.23.118:5178 | ✓ 1581ms | 否 | ✓ 1284ms | ✓ 1618ms | ✓ 1266ms | http |
| 103.215.36.88:17633 | ✓ 1045ms | ✓ 1802ms | 否 | ✓ 1512ms | 否 | http |
| 147.45.72.6:3128 | ✓ 804ms | 否 | ✓ 1710ms | 否 | ✓ 1518ms | http |
| 81.70.169.194:80 | ✓ 985ms | 否 | ✓ 1599ms | ✓ 1405ms | 否 | http |
| 103.215.36.88:15917 | ✓ 1098ms | ✓ 1473ms | 否 | ✓ 1697ms | ✓ 1072ms | http |
| 74.208.234.198:443 | ✓ 1517ms | 否 | ✓ 1821ms | ✓ 1434ms | 否 | http |
| 154.90.48.209:9090 | ✓ 1776ms | 否 | ✓ 1225ms | ✓ 1641ms | ✓ 1268ms | http |
| 103.139.138.194:3128 | ✓ 1815ms | 否 | ✓ 1597ms | ✓ 1814ms | ✓ 1223ms | http |
| 200.125.171.254:999 | ✓ 1400ms | ✓ 1667ms | ✓ 1319ms | 否 | ✓ 1420ms | http |
| 14.56.177.44:3128 | ✓ 1735ms | 否 | ✓ 1160ms | ✓ 936ms | ✓ 781ms | http |
| 103.215.36.88:16777 | ✓ 1013ms | 否 | 否 | ✓ 1315ms | ✓ 1010ms | http |
| 185.191.236.162:3128 | ✓ 924ms | ✓ 1885ms | ✓ 746ms | 否 | ✓ 1249ms | http |
| 61.72.221.94:3128 | ✓ 1660ms | 否 | ✓ 981ms | 否 | ✓ 1046ms | http |
| 178.236.245.17:3128 | ✓ 1194ms | 否 | ✓ 1530ms | 否 | ✓ 1827ms | http |
| 178.236.245.59:3128 | ✓ 1044ms | 否 | ✓ 1651ms | 否 | ✓ 1907ms | http |
| 194.59.204.87:9080 | ✓ 672ms | ✓ 1857ms | ✓ 793ms | 否 | 否 | http |
| 107.174.80.186:3128 | ✓ 905ms | 否 | ✓ 997ms | ✓ 1360ms | 否 | http |

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
