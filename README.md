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

最后更新：2026-03-14 09:35:10 UTC（2026-03-14 17:35:10 UTC+8）

**代理总数：70**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 216.180.127.45:1080 | ✓ 770ms | 否 | ✓ 1188ms | ✓ 1204ms | ✓ 921ms | http |
| 205.209.118.30:3138 | ✓ 771ms | 否 | 否 | ✓ 1391ms | ✓ 1300ms | http |
| 85.198.96.242:3128 | ✓ 864ms | 否 | 否 | ✓ 1916ms | ✓ 1425ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1325ms | ✓ 1183ms | ✓ 931ms | http |
| 45.167.124.52:8080 | ✓ 1680ms | 否 | 否 | ✓ 1839ms | ✓ 1742ms | http |
| 38.145.203.135:8443 | 否 | ✓ 753ms | ✓ 903ms | ✓ 791ms | ✓ 701ms | http |
| 120.232.242.119:22222 | 否 | 否 | ✓ 1011ms | ✓ 1064ms | ✓ 867ms | http |
| 35.225.22.61:80 | ✓ 470ms | 否 | ✓ 911ms | ✓ 1420ms | ✓ 1106ms | http |
| 210.77.29.245:7890 | ✓ 911ms | ✓ 996ms | ✓ 1040ms | ✓ 1057ms | ✓ 804ms | http |
| 120.92.212.16:7890 | ✓ 1093ms | ✓ 1198ms | ✓ 1227ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 957ms | 否 | ✓ 1206ms | ✓ 1160ms | 否 | http |
| 162.243.149.86:31028 | ✓ 1063ms | ✓ 869ms | ✓ 1141ms | ✓ 879ms | ✓ 606ms | http |
| 202.155.12.161:443 | ✓ 1297ms | 否 | ✓ 676ms | ✓ 1134ms | ✓ 961ms | http |
| 150.230.249.50:1080 | ✓ 1330ms | 否 | ✓ 1786ms | ✓ 1892ms | ✓ 707ms | http |
| 81.70.169.194:80 | 否 | 否 | ✓ 1627ms | ✓ 1518ms | ✓ 1004ms | http |
| 101.43.255.96:80 | ✓ 937ms | 否 | 否 | ✓ 1279ms | ✓ 1360ms | http |
| 183.249.5.109:22222 | ✓ 697ms | ✓ 899ms | ✓ 724ms | 否 | ✓ 718ms | http |
| 103.113.70.189:1081 | ✓ 377ms | 否 | 否 | ✓ 1510ms | ✓ 974ms | http |
| 141.98.197.133:18791 | ✓ 1618ms | ✓ 991ms | ✓ 1303ms | ✓ 867ms | ✓ 765ms | http |
| 117.159.239.46:22222 | ✓ 1022ms | ✓ 1022ms | ✓ 1019ms | ✓ 1081ms | ✓ 824ms | http |
| 45.88.0.98:3128 | ✓ 1252ms | 否 | ✓ 649ms | ✓ 1570ms | ✓ 1711ms | http |
| 213.220.62.62:3128 | ✓ 1290ms | 否 | ✓ 1811ms | 否 | ✓ 1240ms | http |
| 45.88.0.113:3128 | ✓ 1966ms | 否 | ✓ 1170ms | ✓ 1588ms | 否 | http |
| 45.88.0.115:3128 | ✓ 654ms | ✓ 1965ms | ✓ 760ms | ✓ 1614ms | ✓ 1277ms | http |
| 45.88.0.99:3128 | ✓ 691ms | 否 | ✓ 681ms | ✓ 1605ms | ✓ 1317ms | http |
| 45.88.0.117:3128 | ✓ 648ms | 否 | ✓ 1160ms | ✓ 1602ms | 否 | http |
| 45.88.0.111:3128 | ✓ 1808ms | 否 | ✓ 654ms | ✓ 1594ms | ✓ 1227ms | http |
| 45.88.0.116:3128 | ✓ 657ms | 否 | ✓ 1162ms | ✓ 1604ms | 否 | http |
| 123.57.0.163:8888 | ✓ 1859ms | ✓ 1976ms | ✓ 1824ms | 否 | ✓ 1374ms | http |
| 46.246.1.106:3128 | ✓ 1535ms | ✓ 1773ms | ✓ 1965ms | 否 | 否 | http |
| 89.251.9.11:3128 | ✓ 942ms | ✓ 1874ms | 否 | ✓ 1693ms | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 834ms | ✓ 917ms | ✓ 662ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1151ms | ✓ 1496ms | ✓ 1610ms | http |
| 86.53.183.16:1080 | ✓ 1665ms | 否 | ✓ 1875ms | 否 | ✓ 1663ms | http |
| 45.88.0.114:3128 | ✓ 1498ms | 否 | ✓ 1945ms | 否 | ✓ 1840ms | http |
| 116.80.49.165:3172 | ✓ 1469ms | 否 | ✓ 1453ms | 否 | ✓ 1607ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1242ms | ✓ 1253ms | 否 | ✓ 1748ms | http |
| 120.238.159.229:22222 | ✓ 828ms | ✓ 1256ms | ✓ 863ms | 否 | 否 | http |
| 120.240.29.51:22222 | ✓ 1720ms | 否 | ✓ 881ms | ✓ 1081ms | 否 | http |
| 113.59.32.162:22222 | 否 | ✓ 1333ms | ✓ 943ms | ✓ 1210ms | ✓ 946ms | http |
| 120.92.211.211:7890 | ✓ 990ms | ✓ 1721ms | 否 | ✓ 1430ms | ✓ 1675ms | http |
| 106.117.208.101:7890 | ✓ 1014ms | 否 | ✓ 997ms | 否 | ✓ 1150ms | http |
| 103.82.93.98:3128 | ✓ 787ms | 否 | ✓ 1236ms | ✓ 1370ms | ✓ 998ms | http |
| 120.240.35.173:22222 | 否 | ✓ 1163ms | ✓ 945ms | ✓ 1933ms | ✓ 953ms | http |
| 183.249.5.105:22222 | ✓ 787ms | ✓ 1072ms | ✓ 716ms | ✓ 929ms | ✓ 723ms | http |
| 117.159.239.51:22222 | ✓ 806ms | ✓ 997ms | ✓ 786ms | ✓ 1049ms | ✓ 817ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1784ms | ✓ 510ms | ✓ 802ms | ✓ 1330ms | http |
| 43.167.227.161:1080 | ✓ 1583ms | ✓ 1238ms | 否 | ✓ 728ms | ✓ 613ms | http |
| 217.160.162.25:8888 | ✓ 748ms | 否 | ✓ 966ms | ✓ 1886ms | ✓ 1428ms | http |
| 45.136.198.40:3128 | ✓ 1029ms | ✓ 1829ms | 否 | 否 | ✓ 1791ms | http |
| 128.199.120.45:9090 | 否 | 否 | ✓ 1692ms | ✓ 1913ms | ✓ 1156ms | http |
| 186.148.180.46:999 | ✓ 785ms | 否 | ✓ 1508ms | ✓ 1936ms | ✓ 1674ms | http |
| 223.16.170.103:3128 | ✓ 1334ms | 否 | ✓ 986ms | 否 | ✓ 1028ms | http |
| 120.238.159.228:22222 | ✓ 875ms | ✓ 1194ms | ✓ 903ms | ✓ 1138ms | ✓ 902ms | http |
| 207.254.71.62:8088 | ✓ 864ms | ✓ 1864ms | ✓ 1862ms | 否 | ✓ 1929ms | http |
| 45.136.131.42:8447 | ✓ 323ms | 否 | ✓ 1403ms | ✓ 1264ms | ✓ 533ms | http |
| 222.184.48.248:22222 | ✓ 1995ms | 否 | ✓ 1547ms | ✓ 1128ms | ✓ 832ms | http |
| 222.184.48.252:22222 | ✓ 914ms | ✓ 1591ms | ✓ 1972ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1401ms | ✓ 1016ms | ✓ 706ms | ✓ 900ms | ✓ 730ms | http |
| 45.136.131.39:8443 | ✓ 378ms | ✓ 1825ms | ✓ 581ms | ✓ 777ms | ✓ 592ms | http |
| 165.227.5.10:8888 | ✓ 408ms | 否 | ✓ 793ms | ✓ 676ms | 否 | http |
| 162.240.154.26:3128 | 否 | ✓ 1790ms | ✓ 774ms | ✓ 1715ms | 否 | http |
| 117.159.239.54:22222 | ✓ 811ms | ✓ 984ms | ✓ 831ms | ✓ 1199ms | 否 | http |
| 62.60.177.204:34094 | ✓ 290ms | 否 | ✓ 972ms | ✓ 1097ms | ✓ 1221ms | http |
| 180.127.149.244:1080 | ✓ 980ms | ✓ 1109ms | ✓ 1387ms | ✓ 1946ms | ✓ 1497ms | http |
| 101.43.127.100:8877 | ✓ 806ms | ✓ 1524ms | ✓ 798ms | ✓ 1634ms | 否 | http |
| 223.16.170.103:80 | ✓ 1031ms | 否 | ✓ 1281ms | ✓ 1004ms | ✓ 1041ms | http |
| 183.249.5.117:22222 | ✓ 882ms | ✓ 830ms | ✓ 729ms | ✓ 905ms | ✓ 959ms | http |
| 117.159.239.52:22222 | ✓ 885ms | ✓ 974ms | ✓ 754ms | ✓ 1189ms | ✓ 827ms | http |
| 111.79.111.126:3128 | ✓ 1941ms | ✓ 1624ms | ✓ 1949ms | 否 | ✓ 1943ms | http |

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
