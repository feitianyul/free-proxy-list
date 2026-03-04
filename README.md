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

最后更新：2026-03-04 16:42:16 UTC（2026-03-05 00:42:16 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 120.92.212.16:8890 | ✓ 955ms | ✓ 1223ms | ✓ 977ms | 否 | 否 | http |
| 217.76.245.80:999 | ✓ 1329ms | ✓ 1636ms | ✓ 1545ms | ✓ 1384ms | ✓ 1292ms | http |
| 120.232.242.119:22222 | 否 | 否 | ✓ 983ms | ✓ 1174ms | ✓ 957ms | http |
| 120.92.212.16:7890 | ✓ 1193ms | ✓ 1255ms | 否 | ✓ 1202ms | ✓ 993ms | http |
| 101.43.255.96:80 | ✓ 1030ms | ✓ 1483ms | ✓ 1448ms | ✓ 1454ms | ✓ 1192ms | http |
| 213.111.146.36:26701 | ✓ 939ms | 否 | ✓ 1635ms | 否 | ✓ 1499ms | http |
| 81.70.169.194:80 | ✓ 1365ms | ✓ 1662ms | ✓ 1963ms | ✓ 1889ms | ✓ 1119ms | http |
| 91.193.240.157:9877 | ✓ 1589ms | 否 | ✓ 1480ms | 否 | ✓ 1992ms | http |
| 117.159.239.51:22222 | ✓ 815ms | ✓ 1013ms | ✓ 838ms | ✓ 1104ms | ✓ 847ms | http |
| 120.240.29.168:22222 | 否 | ✓ 1364ms | ✓ 1069ms | ✓ 1299ms | ✓ 1049ms | http |
| 117.159.239.44:22222 | ✓ 824ms | ✓ 1027ms | ✓ 786ms | ✓ 1099ms | ✓ 838ms | http |
| 120.240.35.161:22222 | ✓ 1005ms | ✓ 1190ms | 否 | ✓ 1239ms | ✓ 996ms | http |
| 120.240.35.178:22222 | ✓ 977ms | ✓ 1181ms | ✓ 1006ms | ✓ 1327ms | ✓ 1743ms | http |
| 120.240.35.175:22222 | ✓ 920ms | ✓ 1215ms | ✓ 862ms | ✓ 1201ms | ✓ 1030ms | http |
| 192.166.82.55:1080 | 否 | 否 | ✓ 1628ms | ✓ 1753ms | ✓ 1587ms | http |
| 103.84.95.54:7890 | ✓ 1126ms | 否 | ✓ 803ms | 否 | ✓ 753ms | http |
| 222.184.48.248:22222 | ✓ 1969ms | 否 | ✓ 1253ms | ✓ 1269ms | 否 | http |
| 35.225.22.61:80 | ✓ 710ms | ✓ 1841ms | ✓ 543ms | 否 | 否 | http |
| 35.234.17.221:8080 | ✓ 812ms | 否 | ✓ 1340ms | ✓ 1106ms | 否 | http |
| 113.59.32.148:22222 | 否 | ✓ 1270ms | ✓ 1138ms | ✓ 1166ms | ✓ 934ms | http |
| 120.240.29.174:22222 | ✓ 1064ms | ✓ 1256ms | ✓ 1122ms | ✓ 1225ms | 否 | http |
| 162.240.154.26:3128 | ✓ 1876ms | 否 | ✓ 773ms | 否 | ✓ 1611ms | http |
| 106.14.203.63:3333 | ✓ 812ms | ✓ 1093ms | ✓ 901ms | ✓ 1100ms | ✓ 867ms | http |
| 101.32.244.83:8080 | ✓ 988ms | ✓ 1767ms | ✓ 941ms | ✓ 1438ms | ✓ 1208ms | http |
| 121.43.196.210:8222 | ✓ 957ms | ✓ 1067ms | ✓ 837ms | ✓ 1125ms | ✓ 936ms | http |
| 121.43.196.213:8222 | ✓ 917ms | ✓ 1041ms | ✓ 905ms | ✓ 1139ms | ✓ 889ms | http |
| 114.55.226.123:10086 | ✓ 1086ms | ✓ 1387ms | ✓ 1015ms | ✓ 1264ms | ✓ 1067ms | http |
| 74.48.78.224:2080 | ✓ 639ms | ✓ 1284ms | 否 | 否 | ✓ 781ms | http |
| 120.240.35.177:22222 | ✓ 1039ms | 否 | 否 | ✓ 1262ms | ✓ 950ms | http |
| 88.80.150.82:8080 | ✓ 950ms | 否 | ✓ 922ms | 否 | ✓ 1561ms | https |
| 120.240.35.160:22222 | 否 | ✓ 1290ms | ✓ 1012ms | ✓ 1233ms | ✓ 1009ms | http |
| 120.198.141.75:22222 | ✓ 960ms | ✓ 1348ms | ✓ 1016ms | ✓ 1274ms | ✓ 1027ms | http |
| 45.136.198.40:3128 | ✓ 1126ms | 否 | ✓ 815ms | 否 | ✓ 1279ms | http |
| 62.113.119.14:8080 | ✓ 887ms | ✓ 1787ms | ✓ 1181ms | ✓ 1675ms | ✓ 1236ms | http |
| 211.171.114.154:3128 | ✓ 789ms | ✓ 1961ms | 否 | ✓ 1388ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1061ms | ✓ 1949ms | ✓ 1885ms | ✓ 1757ms | ✓ 1681ms | http |
| 113.59.32.142:22222 | ✓ 1137ms | ✓ 1319ms | ✓ 916ms | ✓ 1187ms | ✓ 917ms | http |
| 103.82.23.118:5234 | ✓ 1453ms | 否 | ✓ 1304ms | 否 | ✓ 1224ms | http |
| 103.215.36.88:16316 | ✓ 1677ms | ✓ 1202ms | ✓ 1543ms | 否 | ✓ 977ms | http |
| 138.124.53.25:7443 | ✓ 1129ms | 否 | ✓ 1069ms | ✓ 1630ms | 否 | http |
| 172.212.68.37:3128 | ✓ 464ms | 否 | ✓ 1710ms | ✓ 1738ms | ✓ 1201ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1244ms | ✓ 1324ms | 否 | ✓ 1288ms | http |
| 120.240.29.51:22222 | ✓ 1041ms | ✓ 1357ms | ✓ 1062ms | ✓ 1219ms | ✓ 910ms | http |
| 121.230.8.136:1080 | ✓ 1867ms | ✓ 1526ms | ✓ 1099ms | 否 | ✓ 1294ms | http |
| 103.39.51.190:8080 | ✓ 1844ms | 否 | ✓ 1837ms | ✓ 1332ms | ✓ 1355ms | http |
| 120.240.35.174:22222 | ✓ 1012ms | ✓ 1163ms | ✓ 1023ms | ✓ 1209ms | ✓ 1045ms | http |
| 5.75.196.26:40000 | ✓ 1238ms | ✓ 1876ms | ✓ 1859ms | 否 | 否 | http |
| 77.83.203.6:443 | ✓ 1347ms | 否 | ✓ 1945ms | 否 | ✓ 1495ms | http |
| 121.230.8.220:1080 | 否 | ✓ 1192ms | ✓ 1044ms | ✓ 1498ms | ✓ 1243ms | http |
| 117.159.239.52:22222 | ✓ 890ms | ✓ 1052ms | ✓ 813ms | ✓ 1061ms | ✓ 857ms | http |
| 117.159.239.54:22222 | ✓ 858ms | ✓ 1023ms | ✓ 819ms | ✓ 1131ms | ✓ 877ms | http |
| 120.55.163.237:10086 | 否 | 否 | ✓ 1717ms | ✓ 1106ms | ✓ 855ms | http |
| 113.59.32.162:22222 | ✓ 1174ms | ✓ 1329ms | ✓ 1051ms | ✓ 1272ms | ✓ 932ms | http |
| 160.238.65.3:3128 | ✓ 1155ms | ✓ 1745ms | 否 | 否 | ✓ 1708ms | http |
| 160.238.65.2:3128 | ✓ 1150ms | 否 | ✓ 1790ms | 否 | ✓ 1676ms | http |
| 160.238.65.5:3128 | ✓ 1150ms | 否 | ✓ 1788ms | 否 | ✓ 1684ms | http |
| 160.238.65.6:3128 | ✓ 1155ms | ✓ 1933ms | ✓ 1852ms | 否 | ✓ 1704ms | http |
| 160.238.65.9:3128 | ✓ 1148ms | ✓ 1995ms | ✓ 1796ms | 否 | ✓ 1669ms | http |
| 160.238.65.8:3128 | ✓ 1149ms | 否 | ✓ 1791ms | 否 | ✓ 1689ms | http |
| 160.238.65.7:3128 | ✓ 1150ms | 否 | ✓ 1791ms | 否 | ✓ 1677ms | http |
| 160.238.65.4:3128 | ✓ 1150ms | 否 | ✓ 1790ms | 否 | ✓ 1695ms | http |
| 165.227.5.10:8888 | ✓ 158ms | 否 | 否 | ✓ 864ms | ✓ 602ms | http |

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
