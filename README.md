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

最后更新：2026-03-16 14:15:02 UTC（2026-03-16 22:15:02 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 8.209.239.31:30000 | ✓ 452ms | ✓ 1065ms | ✓ 526ms | ✓ 741ms | ✓ 558ms | http |
| 202.155.12.161:443 | ✓ 902ms | 否 | ✓ 888ms | ✓ 1223ms | 否 | http |
| 35.225.22.61:80 | ✓ 662ms | 否 | 否 | ✓ 1236ms | ✓ 1196ms | http |
| 45.140.147.155:1081 | ✓ 1204ms | 否 | ✓ 1224ms | ✓ 1872ms | ✓ 1661ms | http |
| 217.76.245.80:999 | ✓ 1054ms | 否 | ✓ 1194ms | ✓ 1625ms | ✓ 1322ms | http |
| 147.161.210.140:8800 | ✓ 1349ms | 否 | ✓ 893ms | ✓ 1022ms | ✓ 1406ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 265ms | ✓ 949ms | ✓ 977ms | http |
| 115.231.181.40:8128 | ✓ 952ms | 否 | ✓ 1291ms | 否 | ✓ 965ms | http |
| 59.46.216.131:30001 | ✓ 1969ms | 否 | ✓ 1279ms | ✓ 1806ms | 否 | http |
| 45.167.124.52:8080 | ✓ 1032ms | 否 | ✓ 1339ms | ✓ 1729ms | ✓ 1417ms | http |
| 120.92.212.16:7890 | ✓ 1073ms | ✓ 1267ms | ✓ 963ms | 否 | 否 | http |
| 137.220.150.152:6005 | ✓ 1656ms | 否 | ✓ 906ms | 否 | ✓ 1036ms | http |
| 38.34.179.175:8448 | 否 | 否 | ✓ 952ms | ✓ 1129ms | ✓ 1664ms | http |
| 38.34.179.23:8448 | ✓ 429ms | ✓ 1919ms | ✓ 267ms | ✓ 1275ms | ✓ 1009ms | http |
| 20.27.13.35:8561 | ✓ 1021ms | ✓ 1037ms | ✓ 790ms | ✓ 787ms | ✓ 715ms | http |
| 20.210.39.153:8561 | ✓ 1018ms | ✓ 1144ms | ✓ 688ms | ✓ 850ms | ✓ 659ms | http |
| 20.27.14.220:8561 | ✓ 1018ms | ✓ 1139ms | ✓ 694ms | ✓ 796ms | ✓ 731ms | http |
| 1.231.81.166:3128 | ✓ 1033ms | ✓ 1207ms | 否 | ✓ 1227ms | ✓ 1040ms | http |
| 38.34.179.14:8450 | ✓ 804ms | ✓ 1495ms | 否 | 否 | ✓ 960ms | http |
| 137.220.151.110:6005 | ✓ 1646ms | 否 | ✓ 903ms | ✓ 1385ms | ✓ 1290ms | http |
| 38.145.220.15:8450 | 否 | 否 | ✓ 809ms | ✓ 1634ms | ✓ 1861ms | http |
| 113.160.132.26:8080 | ✓ 1872ms | ✓ 1600ms | 否 | ✓ 1894ms | ✓ 1012ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1254ms | 否 | ✓ 1286ms | ✓ 1004ms | http |
| 101.43.127.100:8877 | ✓ 986ms | ✓ 1912ms | ✓ 1212ms | 否 | 否 | http |
| 38.34.179.78:8448 | ✓ 1869ms | 否 | ✓ 798ms | ✓ 801ms | ✓ 528ms | http |
| 116.80.65.85:3172 | ✓ 1950ms | 否 | ✓ 1465ms | ✓ 1795ms | 否 | http |
| 116.80.96.100:3172 | 否 | 否 | ✓ 1475ms | ✓ 1815ms | ✓ 1750ms | http |
| 207.254.71.62:8088 | ✓ 1551ms | 否 | ✓ 1586ms | ✓ 1793ms | ✓ 1810ms | http |
| 62.60.177.204:34094 | ✓ 1252ms | ✓ 1421ms | 否 | ✓ 1203ms | ✓ 977ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 884ms | ✓ 987ms | ✓ 870ms | http |
| 101.32.244.83:8080 | ✓ 1461ms | 否 | ✓ 939ms | ✓ 1191ms | ✓ 1206ms | http |
| 121.43.196.213:8222 | ✓ 1049ms | ✓ 1076ms | ✓ 839ms | ✓ 1152ms | ✓ 888ms | http |
| 121.43.196.210:8222 | ✓ 915ms | ✓ 1061ms | ✓ 949ms | ✓ 1119ms | ✓ 899ms | http |
| 114.55.226.123:10086 | ✓ 1032ms | 否 | ✓ 1127ms | ✓ 1321ms | ✓ 1349ms | http |
| 168.235.110.63:3128 | ✓ 507ms | 否 | ✓ 1853ms | 否 | ✓ 1010ms | http |
| 20.27.15.111:8561 | ✓ 563ms | 否 | ✓ 592ms | ✓ 842ms | ✓ 603ms | http |
| 38.34.179.60:8450 | ✓ 635ms | ✓ 704ms | ✓ 594ms | ✓ 1805ms | ✓ 610ms | http |
| 123.57.0.163:8888 | ✓ 1444ms | 否 | 否 | ✓ 1907ms | ✓ 1461ms | http |
| 106.117.208.101:7890 | ✓ 1236ms | ✓ 1464ms | ✓ 987ms | 否 | ✓ 1241ms | http |
| 20.78.118.91:8561 | ✓ 470ms | ✓ 1892ms | ✓ 479ms | ✓ 745ms | ✓ 614ms | http |
| 45.140.147.155:1082 | ✓ 651ms | ✓ 1718ms | 否 | ✓ 1823ms | ✓ 1241ms | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1627ms | ✓ 1583ms | ✓ 1373ms | http |
| 88.80.150.82:8080 | ✓ 1035ms | ✓ 1930ms | 否 | 否 | ✓ 1891ms | https |
| 45.136.131.54:8448 | ✓ 534ms | ✓ 1270ms | 否 | ✓ 822ms | ✓ 764ms | http |
| 38.145.220.49:8443 | ✓ 551ms | ✓ 1571ms | ✓ 1409ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1846ms | 否 | 否 | ✓ 1533ms | ✓ 1414ms | http |
| 45.136.131.36:8450 | ✓ 829ms | ✓ 1662ms | ✓ 1500ms | ✓ 1535ms | 否 | http |
| 14.225.212.37:7890 | ✓ 828ms | 否 | ✓ 861ms | ✓ 1136ms | ✓ 883ms | http |
| 20.78.26.206:8561 | ✓ 949ms | ✓ 1594ms | ✓ 474ms | ✓ 747ms | ✓ 672ms | http |
| 221.122.91.36:11273 | 否 | 否 | ✓ 901ms | ✓ 1237ms | ✓ 944ms | http |
| 221.122.91.36:11195 | 否 | 否 | ✓ 877ms | ✓ 1179ms | ✓ 961ms | http |
| 38.34.179.12:8443 | 否 | 否 | ✓ 738ms | ✓ 721ms | ✓ 611ms | http |
| 133.242.138.34:8100 | ✓ 1169ms | 否 | 否 | ✓ 1362ms | ✓ 1738ms | http |
| 149.50.116.240:1080 | ✓ 1415ms | 否 | ✓ 1221ms | 否 | ✓ 1399ms | http |
| 159.223.42.219:3128 | ✓ 865ms | 否 | ✓ 1007ms | ✓ 1296ms | ✓ 1264ms | http |
| 38.34.179.203:8450 | 否 | 否 | ✓ 467ms | ✓ 699ms | ✓ 530ms | http |
| 16.78.119.130:443 | 否 | ✓ 1942ms | ✓ 1428ms | 否 | ✓ 1521ms | http |
| 45.136.131.59:8450 | 否 | 否 | ✓ 106ms | ✓ 706ms | ✓ 1681ms | http |
| 103.67.46.225:3125 | ✓ 1845ms | 否 | 否 | ✓ 1714ms | ✓ 1594ms | http |
| 61.52.131.172:8443 | ✓ 932ms | ✓ 1147ms | ✓ 1036ms | ✓ 1150ms | ✓ 984ms | http |
| 86.53.183.16:1080 | ✓ 1126ms | ✓ 1546ms | ✓ 1475ms | 否 | ✓ 1990ms | http |
| 38.34.179.213:8450 | ✓ 511ms | ✓ 1261ms | ✓ 285ms | ✓ 648ms | ✓ 540ms | http |
| 45.136.198.40:3128 | ✓ 1418ms | ✓ 1704ms | ✓ 1094ms | 否 | 否 | http |
| 213.219.214.45:443 | ✓ 898ms | 否 | ✓ 1512ms | 否 | ✓ 1743ms | http |
| 172.212.68.37:3128 | ✓ 1309ms | 否 | ✓ 1082ms | ✓ 1635ms | ✓ 1675ms | http |

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
